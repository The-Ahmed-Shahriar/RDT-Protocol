// 
// rdt_packet.go
// 
// Implements the basic RDT packet data structure's functionalities.
// This file implements the following:
// 
// Functions:
// 	func Packet(ptype, seqnum int, msg string) (*RDTPacket,error)
// 

package rdt




func Packet(ptype, seqnum int, msg string) (*RDTPacket,error) {

	// Check integer sizes
	if uint32(ptype) > ^uint32(0) {
		return nil,INVALID_PKT_INT
	}

	if uint32(seqnum) > ^uint32(0) {
		return nil,INVALID_PKT_INT
	}

	// Check message data length
	if len(msg) > DATA_SIZE {
		return nil,INVALID_PKT_STR
	}

	packet := &RDTPacket{ptype, seqnum, len(msg), msg}
	return packet,nil
}


func (packet *RDTPacket) String() (string,error) {

	lengthStr, err := Itoa32(packet.length)
	if err != nil {
		return "",err
	}

	return lengthStr,nil
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

