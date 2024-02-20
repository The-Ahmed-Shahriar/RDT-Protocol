// 
// rdt_conn.go
// 
// This file contains the part of the RDT wrapper that implements the net.Conn interface
// above the provided connection. I.e., *RDTConn is a net.Conn that contains a net.Conn
// (both aggregated and composed). For RDT specific functionalities, see rdt.go
// 
// The justification for this design is to
// 
// (i)  ensure that RDT can be applied over any other implemented connection protocol
//      (even another RDT wrapped protocol), and
// 
// (ii) can still be utilized almost the same as any other connection would
// 
// 
// This following are defined in this file:
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
// Constructors:
// 	func DialRDT(network string, laddr, raddr *net.UDPAddr) (*RDTConn,error)
// 	func ListenRDT(network string, laddr *net.UDPAddr) (*RDTConn, error)
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
		n, err := rdtconn.Conn.Read(data[0:])
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
// Implements net.Conn.Write(); DO NOT USE - unsafe.
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

		n, err := rdtconn.Conn.Write([]byte(str))
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
	_, err = rdtconn.Conn.Write([]byte(str))

	return N,err
}




// (s) DialRDT()
// 
// @param network: network type (recommended type: udp, udp4, udp6)
// @param address: connection remote address
// 
// @return (rdtconn,err): New outgoing RDT connection
// 
// Mimics net.Dial() - establishes a new outgoing RDT connection.
// 
func DialRDT(network, address string) (*RDTConn,error) {

	// Simply establish an outgoing connection
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil,err
	}

	return &RDTConn{conn},nil
}


// (s) ListenRDT()
// 
// @param network: network type (recommended type: udp, udp4, udp6)
// @param laddr: Local address
// 
// @return (rdtconn,err) - New incoming RDT connection
// 
// Mimics net.Listen() - establishes a new incoming RDT connection. Note the difference
// however; the function blocks until an incoming connection is found.
// 
func ListenRDT(network, address string) (*RDTConn, error) {

	var conn net.Conn
	switch network {

	case "udp","udp4","udp6":
		// Use UDP interface; this is not a modular solution for this, unfortunately.
		// And to be honest, this is the best I can come up with atm
		addr, err := net.ResolveUDPAddr(network, address)
		if err != nil {
			return nil,err
		}

		conn, err = net.ListenUDP(network, addr)
		if err != nil {
			return nil,err
		}

	default:
		// Start listening for RDT connection requests
		listener, err := net.Listen(network, address)
		if err != nil {
			return nil,err
		}
		defer listener.Close()

		// Accept connection request (blocking!)
		conn, err = listener.Accept()
		if err != nil {
			return nil,err
		}
	}

	return &RDTConn{conn},nil
}




// (m,i,x) Close()
// 
// @return err
// 
// Implements net.Conn.Close().
// 
func (rdtconn *RDTConn) Close() error {
	return rdtconn.Conn.Close()
}


// (m,i) LocalAddr()
// 
// @return addr
// 
// Implements net.Conn.LocalAddr().
// 
func (rdtconn *RDTConn) LocalAddr() net.Addr {
	return rdtconn.Conn.LocalAddr()
}


// (m,i) RemoteAddr()
// 
// @return addr
// 
// Implements net.Conn.RemoteAddr().
// 
func (rdtconn *RDTConn) RemoteAddr() net.Addr {
	return rdtconn.Conn.RemoteAddr()
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
	return rdtconn.Conn.SetDeadline(t)
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
	return rdtconn.Conn.SetReadDeadline(t)
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
	return rdtconn.Conn.SetWriteDeadline(t)
}

