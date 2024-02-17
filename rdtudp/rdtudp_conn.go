// 
// This file implements the UDP protocol connection extended to
// support reliable data transfer (rdt). This connection struct extends
// from the standard library's net.Conn interface, just like the
// net.UDPConn and net.TCPConn structs.
// 
// This following are defined in this file:
// 
// Types:
// 	type RDTConn struct
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
// Methods:
// 	func (rdtconn *RDTConn) ReadFromRDT(b []byte) (int,*net.UDPAddr,error)
// 	func (rdtconn *RDTConn) WriteToRDT(b []byte, addr *net.UDPAddr) (int,error)
// 	func (rdtconn *RDTConn) WriteACK() error
// 	func (rdtconn *RDTConn) WriteEOT() error
// 
// Functions:
// 	func DialRDT(network string, laddr, raddr *net.UDPAddr) (*RDTConn,error)
// 	func ListenRDT(network string, laddr *net.UDPAddr) (*RDTConn, error)
// 


package rdtudp

import (
	"net"
	"time"
)




// (m,i) Read()
func (rdtconn *RDTConn) Read(b []byte) (int,error) {

	n, err := 


// (m,i) Write()




// (m) ReadFromRDT()


// (m) WriteToRDT()




// (m) WriteACK()


// (m) WriteEOT()




// DialRDT()
// 
// @param network: Network type (udp, udp4, udp6)
// @param laddr: Local address
// @param raddr: Remote address
// 
// @return (rdtconn,err): New outgoing RDT UDP connection
// 
// Similar to net.DialUDP() - establishes a new outgoing RDT UDP connection.
// 
func DialRDT(network string, laddr, raddr *net.UDPAddr) (*RDTConn,error) {

	conn, err := DialUDP(network, laddr, raddr)
	if err != nil {
		return nil,err
	}

	return &RDTConn{conn},nil
}


// ListenRDT()
// 
// @param network: Network type (udp, udp4, udp6)
// @param laddr: Local address
// 
// @return (rdtconn,err) - New incoming RDT UDP connection
// 
// Similar to net.ListenUDP() - establishes a new incoming RDT UDP connection.
// 
func ListenRDT(network string, laddr *net.UDPAddr) (*RDTConn, error) {

	conn, err := ListenUDP(network, laddr)
	if err != nil {
		return nil,err
	}

	return &RDTConn{conn},nil
}




// (m,i) Close()
// 
// @return err
// 
// Implements net.Conn.Close()
// 
func (rdtconn *RDTConn) Close() error {
	return rdtconn.Conn.Close()
}


// (m,i) LocalAddr()
// 
// @return addr
// 
// Implements net.Conn.LocalAddr()
// 
func (rdtconn *RDTConn) LocalAddr() net.Addr {
	return rdtconn.Conn.LocalAddr()
}


// (m,i) RemoteAddr()
// 
// @return addr
// 
// Implements net.Conn.RemoteAddr()
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
// Implements net.Conn.SetDeadline()
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
// Implements net.Conn.SetReadDeadline()
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
// Implements net.Conn.SetWriteDeadline()
// 
func (rdtconn *RDTConn) SetWriteDeadline(t time.Time) error {
	return rdtconn.Conn.SetWriteDeadline(t)
}

