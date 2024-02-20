package main

import (
	"a2/rdt"
	"fmt"
	"log"
)




func main() {

	conn, err := rdt.ListenRDT("tcp", "127.0.0.1:9992")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Success!")
}

