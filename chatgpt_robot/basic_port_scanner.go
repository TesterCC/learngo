package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

/*
chat with chatGPT
Q: write a port scanner by go language

A: Code ....
This program takes a single argument, the hostname or IP address of the target machine, and scans all 65535 TCP ports
to determine which ones are open. When a port is found to be open, it will print the address (host and port)
to the console.
The program uses Goroutines and a WaitGroup to efficiently handle multiple concurrent scans.

testcase:
go run basic_port_scanner.go 192.168.170.143
*/

func scanPort(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := host + ":" + strconv.Itoa(port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return
	}
	conn.Close()
	fmt.Println(address, "is open")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run port-scanner.go <host>")
		os.Exit(1)
	}

	host := os.Args[1]
	var wg sync.WaitGroup
	for port := 1; port <= 65535; port++ {
		wg.Add(1)
		go scanPort(host, port, &wg)
	}

	wg.Wait()
	fmt.Println("Scan completed.")
}
