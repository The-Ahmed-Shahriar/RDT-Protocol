#!/bin/bash

# Run script for sender as part of
# Assignment 2
# Computer Networks (CS 456)
# Number of parameters: 5
# Parameter:
# 	$1: <hostname of the network emulator>
# 	$2: <UDP port number used by the emulator to fetch data from the sender>
# 	$3: <UDP port number used by the sender to fetch data from the emulator>
# 	$4: <timeout interval in ms>
# 	$5: <name of the file being transferred>


# For Golang implementation (custom)
go run sender.go $1 $2 $3 $4 "$5"
