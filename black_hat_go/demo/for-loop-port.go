package main

import "fmt"

// P25
var target = "scanme.nmap.org"
func main() {
	for i:= 1; i< 1024; i++ {
		address := fmt.Sprintf(target+":%d", i)
		fmt.Println(address)
	}
}
