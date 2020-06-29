package main

import (
	"fmt"
	"net"
)

// P37

func main() {
	_,err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection successful")
	}
}
