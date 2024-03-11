// 
// rdt_packet.go
// 
// Implements the basic RDT packet data structure's functionalities.
// This file implements the following:
// 
// Constructors:
// 	func ParsePacket(str string) (Packet,error)
// 	func PacketData(seqnum int, msg string) (Packet,error)
// 	func PacketACK(seqnum int) (Packet,error)
// 	func PacketEOT() Packet
// 
// Methods:
// 	func (packet Packet) String() (string,error)
// 	func (packet Packet) Ptype() int
// 	func (packet Packet) Seqnum() int
// 	func (packet Packet) Length() int
// 	func (packet Packet) Data() string
// 
// Conversions:
// 	func Itoa32(num int) (string,error)
// 	func Atoi32(str string) (int,error)
// 

package rdt




func ParsePacket(str string) (Packet,error) {

	// Decode integer parts
	ptype,err := Atoi32(str[0:4])
	if err != nil {
		return nil,err
	}

	seqnum,err := Atoi32(str[4:8])
	if err != nil {
		return nil,err
	}

	length,err := Atoi32(str[8:12])
	if err != nil {
		return nil,err
	}

	// Extract string message with checks
	if length > DATA_SIZE {
		return nil,INVALID_PKT_STR
	}
	msg := str[12:12+length]

	return Packet(ptype,seqnum,msg)
}


func PacketData(seqnum int, msg string) (Packet,error) {

	// Check sequence number size
	if uint32(seqnum) > ^uint32(0) {
		return nil,INVALID_PKT_INT
	}

	// Check message data length
	if len(msg) > DATA_SIZE {
		return nil,INVALID_PKT_STR
	}

	packet := Packet{DATA_PKT,seqnum,len(msg),msg}
	return packet,nil
}


func PacketACK(seqnum int) (Packet,error) {

	// Check sequence number size
	if uint32(seqnum) > ^uint32(0) {
		return nil,INVALID_PKT_INT
	}

	packet := Packet{ACK_PKT,seqnum,0,""}
	return packet,nil
}


func PacketEOT() Packet {
	packet := Packet{EOT_PKT, 0, 0,""}
	return packet
}




func (packet Packet) String() (string,error) {

	// Encode integer parts
	ptype, err := Itoa32(packet.ptype)
	if err != nil {
		return "",err
	}

	seqnum, err := Itoa32(packet.seqnum)
	if err != nil {
		return "",err
	}

	length, err := Itoa32(packet.length)
	if err != nil {
		return "",err
	}

	// Combine into string
	str := ptype + seqnum + length + packet.data

	return str,nil
}




func (packet Packet) Ptype() int {
	return packet.ptype
}


func (packet Packet) Seqnum() int {
	return packet.seqnum
}


func (packet Packet) Length() int {
	return packet.length
}


func (packet Packet) Data() string {
	return packet.data
}




func Itoa32(num int) (string,error) {

	// Check data validity
	if uint32(num) > ^uint32(0) {
		return "",INVALID_PKT_INT
	}

	// Construct 4 byte string
	str := string((num & Q1) >> 24)
	str += string((num & Q2) >> 16)
	str += string((num & Q3) >> 8)
	str += string(num & Q4)

	return str,nil
}


func Atoi32(str string) (int,error) {

	// Check string validity
	if len(str) != 4 {
		return 0, INVALID_PKT_INT
	}

	// Parse through each byte
	num := int(str[0]) << 24
	num += int(str[1]) << 16
	num += int(str[2]) << 8
	num += int(str[3])

	// Check data validity
	if uint32(num) > ^uint32(0) {
		return 0, INVALID_PKT_INT
	}

	return num,nil
}

