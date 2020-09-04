package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Println("I am work in a single goroutine...")
}

func main() {
	go test()  // 让test在goroutine上并发执行

	// main goroutine sleep 1 second
	time.Sleep(time.Second)   // 如果不加阻塞，主goroutine在调度go test()执行前就可能结束了。
	// 在Go语言中，只要主goroutine(main函数)结束，就意味着整个程序已经运行结束了。
}
