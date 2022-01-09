package main

import (
	"fmt"
	"time"
)

/*
27-创建goroutine
https://www.bilibili.com/video/BV1gf4y1r79E?p=27
*/

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine: i=%d\n", i)
		time.Sleep(1 * time.Second)
	}

}

// 主goroutine
func main() {
	// 创建一个goroutine 去执行newTask()流程
	go newTask()

	// 主退从退
	fmt.Println("main goroutine exit")

	//i := 0
	//
	//for {
	//	i++
	//	fmt.Printf("main goroutine: i = %d\n", i)
	//	time.Sleep(1 * time.Second)
	//}

}
