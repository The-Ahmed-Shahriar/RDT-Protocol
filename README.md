# RDT Protocol
Author: Ahmed Shahriar  




### Programming Language
This project uses the Go programming language. The versioning from the servers is listed below:
```
$ go version
go version go1.17.6 linux/amd64
```




### Directory Structure
The directory contains the following source files:
- `sender.sh` and `receiver.sh`
- `sender.go` and `receiver.go`
- `go.mod`
- `README.md`
- `rdt` package directory:
    - `rdt_conn.go`
    - `rdt_defs.go`
    - `rdt.go`
    - `rdt_io.go`
    - `rdt_pkt.go`




### Compiling
No preliminary commands need to be run to compile the code.




### Executing the Program

**Network**:
```
# Based on how you have a network or network emulator emulator setup...
# If you want to run an emulator (e.g., UW nEmulator) then run
$ ./nEmulator <nEmulator data port> <receiver hostname> <receiver data port> <nEmulator ack port> <sender hostname> <sender ack port> <max delay (ms)> <pkt loss probability> <verbose-mode>
```
**Receiver**:
```
$ ./receiver.sh <nEmulator hostname> <nEmulator ack port> <receiver data port> <filename (can be new file)>
```
**Sender**:
```
$ ./sender.sh <nEmulator hostname> <nEmulator data port> <sender ack port> <timeout (ms)> <filename>
```
The port number does not require the preceding colon character, ":", just a valid integer address will suffice.




### Logging
You have the option to enable/disable all logging by modifying the following flag, located in `rdt/rdt_defs.go`, line 76.
```
76      const LOGGING_ON = true
```
Manually disable by setting flag to false:
```
76      const LOGGING_ON = false
```

