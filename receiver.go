package main

import (
	"a2/rdt"
	"fmt"
	"log"
)




func main() {

	conn, err := rdt.ListenRDT("udp", ":9992")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var data [rdt.DATA_SIZE]byte
	_, err = conn.Read(data[0:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("> " + string(data[0:]))
	fmt.Println("Success!")
}

