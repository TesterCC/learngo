package main

import (
	"fmt"
	"learngo/go_concurrency_microservice/ch4/compute"
)

// go run main.go hello.go
func main() {
	//sayHello()

	params := &compute.IntParams{
		P1:1,
		P2:2,
	}
	fmt.Println(params.Add())
}
