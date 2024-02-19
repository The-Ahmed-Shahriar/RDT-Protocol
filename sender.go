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

	fmt.Println("len: ", len(str))

	lenstr, err := pkt.String()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("encoded: ", lenstr)

	length, err := rdt.Atoi32(lenstr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("parsed len: ", length)
}

