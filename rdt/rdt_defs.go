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
// 	const BATCH_SIZE
// 
// 	const PKT_DATA
// 	const PKT_ACK
// 	const PKT_EOT
// 
// 	const SEQNUM_LOG
// 	const ACK_LOG
// 	const ARRIVAL_LOG
// 
// 	const Q1,Q2,Q3,Q4
// 
// 	const LOGGING_ON
// 
// Errors:
// 	var INVALID_PKT_INT
// 	var INVALID_PKT_STR
// 
// Data Structures:
// 	type Packet struct
// 	type RDTConn struct
// 

package rdt

import (
	"errors"
	"net"
	"time"
)



// Define size contants
const (
	INT_SIZE = 4
	DATA_SIZE = 500
	PKT_SIZE = 512

	BATCH_SIZE = 10
)

// Define packet type constants
const (
	ACK_PKT = 0
	DATA_PKT = 1
	EOT_PKT = 2
)

// Define log filenames
const (
	SEQNUM_LOG = "seqnum.log"
	ACK_LOG = "ack.log"
	ARRIVAL_LOG = "arrival.log"
)

// Define byte masks, from MSD to LSD (Big Endian)
const (
	Q1 = 0xFF000000
	Q2 = 0x00FF0000
	Q3 = 0x0000FF00
	Q4 = 0x000000FF
)

// Flag for enabling logging
const LOGGING_ON = true




// Define error types
var (
	INVALID_PKT_INT = errors.New("ERROR: Packet expects 32 bit / 4 byte integer\n")
	INVALID_PKT_STR = errors.New("ERROR: Packet expects no more than 500 bytes of data\n")
)




// Packet struct, following the specified packet format in assignment details
// The total size (512 bytes) for the RDTPacket structures are ensured in 
// the structures implementation
type Packet struct {
	ptype int
	seqnum int
	length int
	data string
}




// Connection wrapper implementing the Reliable Data Transfer (RDT) functionality
// over two UDP connections. The struct also comes with a general use packet buffer
// (can be used different for implementing various read/write type methods).
type RDTConn struct {
	InConn *net.UDPConn
	OutConn *net.UDPConn
	timeout time.Duration
}

