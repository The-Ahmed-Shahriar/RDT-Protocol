// 
// rdt.go
// 
// Consists of RDT specific implementations.
// 
// Methods:
// 	func (rdtconn *RDTConn) SendFile(filename string) error
// 	func (rdtconn *RDTConn) ReceiveFile(filename string) error
// 	func StartRDT(network, inAddress, outAddress string, timeout time.Duration) (*RDTConn, error)
// 


package rdt


import (
	"io"
	"log"
	"maps"
	"os"
	"time"
)




// (m) SendFile()
// 
func (rdtconn *RDTConn) SendFile(filename string) error {

	// Validate and open file
	_, err := os.Stat(filename)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Use io.ReaderFrom impl of RDT
	_, err = rdtconn.ReadFrom(file)
	return err
}


// (m) ReceiveFile()
// 
func (rdtconn *RDTConn) ReceiveFile(filename string) error {

	// Validate and open file
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Use io.WriterTo impl of RDT
	_, err = rdtconn.WriteTo(file)
	return err
}




// (s) StartRDT()
// 
// @param network: network type (compatible: udp, udp4, udp6)
// @param laddr: Local address
// 
// @return (rdtconn,err) - New RDT connection
// 
// Similar to net.Dial() and net.Listen() combined - establishes a new RDT connection.
// 
func StartRDT(network, inAddress, outAddress string, timeout time.Duration) (*RDTConn, error) {

	// Establish incoming/read UDP connection
	inAddr, err := net.ResolveUDPAddr(network, inAddress)
	if err != nil {
		return nil,err
	}

	inConn, err := net.ListenUDP(network, inAddr)
	if err != nil {
		return nil,err
	}

	// Establish outgoing/write UDP connection
	outAddr, err := net.ResolveUDPAddr(network, outAddress)
	if err != nil {
		inConn.Close()
		return nil,err
	}

	outConn, err = net.ListenUDP(network, outAddr)
	if err != nil {
		inConn.Close()
		return nil,err
	}

	// Initialize the connection with a general packet buffer (slice)
	return &RDTConn{ inConn, outConn, timeout },nil
}

