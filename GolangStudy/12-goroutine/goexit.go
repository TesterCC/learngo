package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 调用无参goroutine
	// 用go创建承载一个形参为空，返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer") //后执行

		// 匿名函数
		func() {
			defer fmt.Println("B.defer") //后执行

			// 退出当前goroutine
			runtime.Goexit()

			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}
}

/*
B
B.defer
A
A.defer

ref: https://www.bilibili.com/video/BV1gf4y1r79E?p=27
6:06 在goroutine中退出
*/
