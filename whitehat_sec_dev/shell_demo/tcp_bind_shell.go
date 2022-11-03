package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

/*
P71 TCP server demo, bind shell basic
The demo function: receive client data
use telnet or nc to connect bind shell

usage:
go run tcp_bind_shell.go 0.0.0.0:9999
./tcp_bind_shell 0.0.0.0:9999

nc 192.168.X.X 9999
*/

func main() {
	var addr string
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " <bindAddress>")
		fmt.Println("Example: " + os.Args[1] + " 0.0.0.0:9999")
		os.Exit(1)
	}
	addr = os.Args[1]

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Error connection!", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accepting connection err: ", err)
		}
		go handleConnectionDemo(conn)
	}

}

func handleConnectionDemo(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff[:])
		if err != nil {
			continue
		}
		_, err = conn.Write(buff[:n])
	}

}
