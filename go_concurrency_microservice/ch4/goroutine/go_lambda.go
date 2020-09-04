package main

import (
	"fmt"
	"time"
)

// P89 go语句除了启动命名函数，还可以启动匿名函数

func main() {

	// 匿名函数
	go func(name string) {
		fmt.Println("Hello, " + name)
	}("xuan")
	// 主goroutine阻塞1s
	time.Sleep(time.Second)
}
