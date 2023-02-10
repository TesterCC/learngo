package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)


//  an example of a basic TCP server and client in Go
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	username := getUsername()
	sendMessage(conn, username)
	message := readMessage(conn)
	fmt.Println("Message from server:", message)
}

func getUsername() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	username, _ := reader.ReadString('\n')
	return strings.TrimSpace(username)
}

func readMessage(conn net.Conn) string {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	return strings.TrimSpace(message)
}

func sendMessage(conn net.Conn, message string) {
	fmt.Fprintln(conn, message)
}
