package main

import (
	"fmt"
	"net"
	"time"
)

// P37

func main() {

	var ip = "10.0.4.69"
	var port = "6379"

	// the simplest port scan
	target := fmt.Sprintf("%s:%s",ip,port)
	//_,err := net.Dial("tcp", "scanme.nmap.org:80")
	_,err := net.Dial("tcp", target)
	if err == nil {
		fmt.Println("Connection successful")
	}
	time.Sleep(1200 * time.Microsecond)
}


