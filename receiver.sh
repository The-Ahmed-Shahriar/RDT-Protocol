#!/bin/bash

# Run script for receiver as part of
# Assignment 2
# Computer Networks (CS 456)
# Number of parameters: 4
# Parameter:
# 	$1: <hostname of the network emulator>
# 	$2: <UDP port number used by the emulator to fetch data from the receiver>
# 	$3: <UDP port number used by the receiver to fetch data from the emulator>
# 	$4: <name of the file into which the received data is written>


# For Golang implementation (custom)
go run receiver.go $1 $2 $3 "$4"
