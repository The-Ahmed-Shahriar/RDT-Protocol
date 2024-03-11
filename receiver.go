package main

import (
	"a2/rdt"
	"fmt"
	"os"
)




func main() {

	inaddr, outaddr, filename := GetArgs()

	rdtconn, err := rdt.StartRDT("udp", inaddr, outaddr, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = rdtconn.ReceiveFile(filename)
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




func GetArgs() (string,string,string) {

	// Validate number of args
	if len(os.Args) != 5 {
		fmt.Println("usage: <receiver script> <remote hostname> :<remote data port> :<local data port> <filename>")
		os.Exit(1)
	}

	// Validate port number format

	return os.Args[3], os.Args[1]+os.Args[2], os.Args[4]
}
