package main

import (
	"a2/rdt"
	"fmt"
	"log"
)




func main() {

	str := "Hello World!\n"

	conn, err := rdt.DialRDT("tcp", ":9992")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(str))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success!")
}

