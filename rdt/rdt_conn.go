// 
// rdt_conn.go
// 
// RDT implementation of the net.Conn, io.ReaderFrom, and io.WriterTo interfaces.
// RDT specific functionality can be found in rdt.go
// 
// The following are defined in this file:
// 
// Methods (implements net.Conn interface):
// 	func (rdtconn *RDTConn) Read(b []byte) (int,error)
// 	func (rdtconn *RDTConn) Write(b []byte) (int,error)
// 	func (rdtconn *RDTConn) Close() error
// 	func (rdtconn *RDTConn) LocalAddr() net.Addr
// 	func (rdtconn *RDTConn) RemoteAddr() net.Addr
// 	func (rdtconn *RDTConn) SetDeadline(t time.Time) error
// 	func (rdtconn *RDTConn) SetReadDeadline(t time.Time) error
// 	func (rdtconn *RDTConn) SetWriteDeadline(t time.Time) error
// 


package rdt


import (
	"net"
	"io"
	"time"
)




// (m,i) Read()
// 
// @param b: Byte slice buffer for returning readings
// 
// @return (N,err): Number of bytes read
// 
// Implements net.Conn.Read(); DO NOT USE - unsafe.
// 
func (rdtconn *RDTConn) Read(b []byte) (int,error) {

	N := 0

	// Keep reading packets until EOT received
	for {
		var data [PKT_SIZE]byte

		// Read raw data from underlying connection
		n, err := rdtconn.InConn.Read(data[0:])
		if err != nil && err != io.EOF {
			return N,err
		}

		// Unmarshall the received data into RDT packet
		packet, err := ParsePacket(string(data[0:]))
		if err != nil {
			return N,err
		}

		// Stop reading if EOT packet found
		if packet.IsEOT() {
			break
		}

		// Copy data portion to output slice
		copy(b[N:], packet.Data())
		N += n
	}

	return N,nil
}


// (m,i) Write()
// 
// @param b: Byte slice buffer for writing out
// 
// @return (N,err): Number of bytes written
// 
// Implements net.Conn.Write(); DO NOT USE - not safe.
// 
func (rdtconn *RDTConn) Write(b []byte) (int,error) {

	N := 0
	numpkts := 1+len(b)/DATA_SIZE

	// Break message into DATA_SIZE pieces to packet
	for i := 0; i < numpkts; i++ {

		// Fix message chunk end index
		end := N+DATA_SIZE
		if end > len(b) {
			end = len(b)
		}

		// Marshall data into a packet
		packet, err := Packet(0, i, string(b[N:end]))
		if err != nil {
			return N,err
		}

		// Write raw data to underlying connection
		str, err := packet.String()
		if err != nil {
			return N,err
		}

		n, err := rdtconn.OutConn.Write([]byte(str))
		if err != nil {
			return N,err
		}
		N += n
	}

	// Write EOT packet
	str, err := EOT.String()
	if err != nil {
		return N,err
	}
	_, err = rdtconn.OutConn.Write([]byte(str))

	return N,err
}




// (m,i,x) Close()
// 
// @return err
// 
// Implements net.Conn.Close().
// 
func (rdtconn *RDTConn) Close() error {
	inerr := rdtconn.InConn.Close()
	outerr := rdtconn.OutConn.Close()

	if inerr != nil {
		return inerr
	}
	return outerr
}


// (m,i) LocalAddr()
// 
// @return addr
// 
// Implements net.Conn.LocalAddr().
// 
func (rdtconn *RDTConn) LocalAddr() net.Addr {
	return rdtconn.InConn.LocalAddr()
}


// (m,i) RemoteAddr()
// 
// @return addr
// 
// Implements net.Conn.RemoteAddr().
// 
func (rdtconn *RDTConn) RemoteAddr() net.Addr {
	return rdtconn.OutConn.RemoteAddr()
}


// (m,i) SetDeadline()
// 
// @param t: Deadline set time
// 
// @return err
// 
// Implements net.Conn.SetDeadline().
// 
func (rdtconn *RDTConn) SetDeadline(t time.Time) error {
	err := rdtconn.InConn.SetReadDeadline(t)
	if err != nil {
		return err
	}
	return rdtconn.OutConn.SetWriteDeadline(t)
}


// (m,i) SetReadDeadline()
// 
// @param t: Deadline set time
// 
// @return err
// 
// Implements net.Conn.SetReadDeadline().
// 
func (rdtconn *RDTConn) SetReadDeadline(t time.Time) error {
	return rdtconn.InConn.SetReadDeadline(t)
}


// (m,i) SetWriteDeadline()
// 
// @param t: Deadline set time
// 
// @return err
// 
// Implements net.Conn.SetWriteDeadline().
// 
func (rdtconn *RDTConn) SetWriteDeadline(t time.Time) error {
	return rdtconn.OutConn.SetWriteDeadline(t)
}

