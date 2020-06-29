package main

import (
	"fmt"
	"time"
)

func action() {
	fmt.Println("Test Goroutine")
}

func main() {
	go action()
	// 沉睡2秒  让主协程等待一段时间的子协程，以免主协程执行完后销毁（这时子协程也会被销毁）
	time.Sleep(2 * time.Second)
}
