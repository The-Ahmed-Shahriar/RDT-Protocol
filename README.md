# CS 456 Assignment 2 Solution
Author: Ahmed Shahriar  




### Programming Language
Like my previous submission, I chose to use Golang. My source code is developed and tested on the school's servers. The versioning from the servers is listed below:
```
$ go version
go version go1.17.6 linux/amd64
```




### Submission Directory
After unzipping, the directory should **hopefully** contain **all** of the following files:
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
The commands are the same as those specified in the assignment description.
**Network Emulator**:
```
# Based on how you have nEmulator setup...
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
The port number does not require the preceding colon character, ":", unlike from my Assignment 1 submission. Just a valid integer will suffice. The "Example Execution" from the assignment description page 4 should work just fine, replacing each `arg[0]` with their respective local commands.

If you are experiencing issues running with the different machines, try using either `ubuntu2204-014`, `ubuntu2204-006`, and `ubuntu2204-012`. I've tested the code on these three servers mostly.




### Logging
Similar to the debugging mode from Assignment 1, you also have the option to enable/disable all logging by modifying the following flag, located in `rdt/rdt_defs.go`, line 76.
```
76      const LOGGING_ON = true
```
Manually disable by setting flag to false:
```
76      const LOGGING_ON = false
```

