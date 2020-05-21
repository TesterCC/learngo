package main

import (
	"fmt"
	"time"
)

// P16 use goroutines

func f(){
	fmt.Println("f function")
}

func main() {
	go f()  // use the go keyword before the call to method or function you wish
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}

