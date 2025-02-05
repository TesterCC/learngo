package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

/*
Q: TCP demo
The server listens on port 8080 and accepts incoming connections. For each connection, it calls the handleConnection function, which reads a message from the client, sends a message to the client, and closes the connection.
The client connects to the server on localhost:8080, sends a message to the server, reads a message from the server, and prints it to the console.
 */

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	username := readMessage(conn)
	fmt.Printf("[D] message from client: %s\n", username)
	sendMessage(conn, "Hello, "+username)
	conn.Close()
}

func readMessage(conn net.Conn) string {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	return strings.TrimSpace(message)
}

func sendMessage(conn net.Conn, message string) {
	fmt.Fprintln(conn, message)
}
