package main

import (
	"a2/rdt"
	"fmt"
	"os"
	"strconv"
	"time"
)




func main() {

	inaddr, outaddr, timeout, filename := GetArgs()

	rdtconn, err := rdt.StartRDT("udp", inaddr, outaddr, timeout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = rdtconn.SendFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = rdtconn.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}




func GetArgs() (string,string,time.Duration,string) {

	// Validate number of args
	if len(os.Args) != 6 {
		fmt.Println("usage: <sender script> <remote hostname> <remote data port> <local ack port> <timeout (ms)> <filename>")
		os.Exit(1)
	}

	// Validate port number format

	// Validate timeout range
	t, err := strconv.Atoi(os.Args[4])
	if err != nil || t < 1 {
		fmt.Println("ERROR: Invalid timeout format")
	}
	timeout := time.Duration(t) * time.Millisecond

	return ":"+os.Args[3], os.Args[1]+":"+os.Args[2], timeout, os.Args[5]
}
