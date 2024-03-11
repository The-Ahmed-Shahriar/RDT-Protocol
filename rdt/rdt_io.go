// 
// rdt_io.go
// 
// RDT implementation of the io.ReaderFrom and io.WriterTo interfaces.
// RDT specific functionality can be found in rdt.go
// 
// The following are defined in this file:
// 
// Methods (implements io.ReaderFrom and io.WriterTo interfaces):
// 	func (rdtconn *RDTConn) ReadFrom(r io.Reader) (int64,error)
// 	func (rdtconn *RDTConn) WriteTo(w io.Writer) (int64,error)
// 

package rdt


import (
	"io"
	"log"
	"os"
	"time"
)




// (m,i) ReadFrom()
// 
// @param r: io.Reader compatible object
// 
// @return (N,err): Number of bytes supplied to r
// 
// Implements io.ReaderFrom.ReadFrom()
// 
func (rdtconn *RDTConn) ReadFrom(r io.Reader) (int64,error) {

	var N int64 = 0
	var sendbase int = 0

	var seqlog, acklog *log.Logger


	/* PRE-PROCESSING */

	// Enable logging
	if LOGGING_ON {
		seqfile, err := os.OpenFile(SEQNUM_LOG, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return N,err
		}
		defer seqfile.Close()

		ackfile, err := os.OpenFile(ACK_LOG, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return N,err
		}
		defer seqfile.Close()

		seqlog = log.New(seqfile, "", 0)
		acklog = log.New(ackfile, "", 0)
	}

	// Use EOT procedure before exiting function - kill receiver
	defer func() {

		// Send EOT pkt
		if LOGGING_ON {
			seqlog.Println("EOT")
		}

		str, err := PacketEOT().String()
		if err != nil {
			return
		}

		_, err = rdtconn.OutConn.Write([]byte(str))
		if err != nil {
			return
		}

		// Await EOT acknowledgement
		for {
			var data [PKT_SIZE]byte

			_, err = rdtconn.InConn.Read(data[0:])
			if err != nil {
				return
			}

			pkt, err := ParsePacket(string(data[0:]))
			if err != nil {
				return
			} else if pkt.Ptype() == EOT_PKT {
				break
			} else if LOGGING_ON && pkt.Ptype() == ACK_PKT {
				acklog.Println(pkt.Seqnum())
			}
		}

		if LOGGING_ON {
			acklog.Println("EOT")
		}
	}()

	// Initialize pkt buffer (batch)
	batch := make(map[int]Packet)


	/* ACK HANDLING GOROUTINE */

	// Define go channels for synchronization
	timerOn := make(chan bool, 1)
	processing := make(chan bool, 1)
	halted := make(chan bool)

	// Immediately start goroutine that always reads for responses
	go func() {

		// Reset any deadline restrictions
		defer rdtconn.SetReadDeadline(time.Time{})

		for {

			var data [PKT_SIZE]byte


			// Seek authority to start reading from socket
			select {
			case <-timerOn:
				break
			case <-halted:
				return
			}

			// Always read for responses; Preempt after 1ms
			rdtconn.SetReadDeadline(time.Now().Add(time.Millisecond))
			_, err := rdtconn.InConn.Read(data[0:])

			// Handle read error
			if err != nil {
				continue
			}

			// Block until timer started or halt triggered
			select {
			case <-timerOn:
				break
			case <-halted:
				return
			}

			/* Handle ACK */
			processing <- true

			ack, err := ParsePacket(string(data[0:]))
			if err != nil {
				<-processing
				continue
			}
			if ack.Ptype() != ACK_PKT {
				<-processing
				return
			}

			// Ignore if acknowdgement's data pkt not in current batch
			_, prs := batch[ack.Seqnum()]
			if !prs {
				<-processing
				continue
			}

			// Remove the corresponding pkt from batch
			if LOGGING_ON {
				acklog.Println(ack.Seqnum())
			}
			delete(batch, ack.Seqnum())
			<-processing
		}
	}()


	/* START FSM */
	readerEOF := false
	for {

		// Load batch with new pkts
		for !readerEOF && len(batch) < BATCH_SIZE {

			var data [DATA_SIZE]byte

			// Fetch next 500 bytes from source
			n, err := r.Read(data[0:])
			if err == io.EOF {
				readerEOF = true
				break
			} else if err != nil {
				halted <- true
				return N,err
			}

			// Construct and buffer new pkt
			batch[sendbase], err = PacketData(sendbase, string(data[0:n]))
			sendbase++
			N += int64(batch[sendbase].Length())
		}

		// Halt if source depleted and all pkts acknowledged
		if readerEOF && len(batch) == 0 {
			break
		}

		// Send batch
		for _, pkt := range batch {

			if LOGGING_ON {
				seqlog.Println(pkt.Seqnum())
			}

			str, err := pkt.String()
			if err != nil {
				halted <- true
				return N,err
			}

			_, err = rdtconn.OutConn.Write([]byte(str))
			if err != nil {
				halted <- true
				return N,err
			}
		}

		// Run timer on duration
		timer := time.NewTimer(rdtconn.timeout)

		// Branch reset point
		timerCheckPoint:

		// Keep ACK handler functioning while timer running and all ACKs not received
		select {
		case <-timer.C:
		default:
			if len(batch) > 0 {

				if len(timerOn) == 0 {
					timerOn <- true
				}
				goto timerCheckPoint

			} else {
				timer.Stop()
			}
		}

		// Sync here, in case response is still being processed
		for len(processing) > 0 {}
	}

	halted <- true
	return N,nil
}




// (m,i) WriteTo()
// 
// @param w: io.Writer compatible object
// 
// @return (N,err): Number of bytes acquired from w
// 
// Implements io.WriterTo.WriteTo()
// 
func (rdtconn *RDTConn) WriteTo(w io.Writer) (int64,error) {

	var N int64 = 0
	var rcvbase int = 0
	var databuff map[int]string
	var arvlog *log.Logger


	/* PRE-PROCESSING */

	// Enable logging
	if LOGGING_ON {
		arvfile, err := os.OpenFile(ARRIVAL_LOG, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return N,err
		}

		defer arvfile.Close()
		arvlog = log.New(arvfile, "", 0)
	}

	// Send EOT pkt before exiting function - kill sender
	defer func() {
		if LOGGING_ON {
			arvlog.Println("EOT")
		}

		str, err := PacketEOT().String()
		if err != nil {
			return
		}

		rdtconn.OutConn.Write([]byte(str))
	}()

	// Initialize data buffer
	databuff = make(map[int]string)


	/* START FSM */
	for {

		var pktbin [PKT_SIZE]byte

		// Read next incoming packet (blocking call)
		_, err := rdtconn.InConn.Read(pktbin[0:])
		if err != nil {
			return N,err
		}
		pkt, err := ParsePacket(string(pktbin[0:]))
		if err != nil {
			return N,err
		}

		// Handle EOT case
		if pkt.Ptype() == EOT_PKT {
			break
		}

		// Send mandatory ACK
		if LOGGING_ON {
			arvlog.Println(pkt.Seqnum())
		}
		ack, err := PacketACK(pkt.Seqnum())
		if err != nil {
			return N,err
		}
		str, err := ack.String()
		if err != nil {
			return N,err
		}
		_, err = rdtconn.OutConn.Write([]byte(str))
		if err != nil {
			return N,err
		}

		// Skip processing if pkt already buffered
		_, prs := databuff[pkt.Seqnum()]
		if prs {
			continue
		}

		// Store new data in buffer
		databuff[pkt.Seqnum()] = pkt.Data()
		N += int64(pkt.Length())

		// Update rcvbase, migrating any contiguous data onto file
		for {

			// Check prescence of rcvbase packet
			str, prs := databuff[rcvbase]
			if !prs {
				break
			}

			// Remove from buffer and update state
			delete(databuff, rcvbase)
			rcvbase++

			// Write data to destination
			_, err = w.Write([]byte(str))
			if err != nil {
				return N,err
			}
		}
	}
	return N,nil
}

