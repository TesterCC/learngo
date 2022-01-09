package main

import (
	"fmt"
	"time"
)

// 调用有参goroutine

func main() {

	//
	go func(a int, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(10, 20)

	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
