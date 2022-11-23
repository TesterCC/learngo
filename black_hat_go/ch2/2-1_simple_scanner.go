package main

import (
	"fmt"
	"net"
	"time"
)

// P24  the simplest port scan - A basic port scanner that scans only one port
// test target: scanme.nmap.org

func main() {

	//var ip = "10.0.4.69"
	//var port = "6379"

	// cat target str
	//target := fmt.Sprintf("%s:%s",ip,port)

	//_,err := net.Dial("tcp", target)
	_,err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection successful")
	}
	time.Sleep(1200 * time.Microsecond)
}
