// 
// rdt_defs.go
// 
// Sets the defintions and resources for the rdt package.
// This file implements the following:
// 
// Constants:
// 	const INT_SIZE
// 	const DATA_SIZE
// 	const PKT_SIZE
// 	var EOT *RDTPacket
// 
// Errors:
// 	var INVALID_PKT_INT
// 	var INVALID_PKT_STR
// 
// Data Structures:
// 	type RDTPacket struct
// 	type RDTConn struct
// 

package rdt

import (
	"errors"
	"net"
)



// Define package contants
const (
	INT_SIZE = 4
	DATA_SIZE = 500
	PKT_SIZE = 512
)


// Define byte masks, from MSD to LSD (Big Endian)
const (
	Q1 = 0xFF000000
	Q2 = 0x00FF0000
	Q3 = 0x0000FF00
	Q4 = 0x000000FF
)


// Define error types
var (
	INVALID_PKT_INT = errors.New("ERROR: Packet expects 32 bit / 4 byte integer\n")
	INVALID_PKT_STR = errors.New("ERROR: Packet expects no more than 500 bytes of data\n")
)


// Define the default EOT packet
var EOT *RDTPacket = &RDTPacket{ 2, 0, 0, "" }




// Packet struct, following the specified packet format in assignment details
// The total size (512 bytes) for the RDTPacket structures are ensured in 
// the structures implementation
type RDTPacket struct {
	ptype int
	seqnum int
	length int
	data string
}




// Connection wrapper implementing the Reliable Data Transfer (RDT) functionality
// over a generic connection type, net.Conn
type RDTConn struct {
	Conn net.Conn
}

