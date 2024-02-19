package main

import (
	"a2/rdt"
	"fmt"
	"os"
)

func main() {

	str := "Hello World!\n"

	pkt, err := rdt.Packet(1,1,str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pktstr, err := pkt.String()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pkt, err = rdt.ParsePacket(pktstr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("data: ", pkt.Data())
}

