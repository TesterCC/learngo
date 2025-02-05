package main

import (
	"fmt"
	"time"
)

/*
29-channel有缓冲与无缓冲同步问题
ref:
https://www.bilibili.com/video/BV1gf4y1r79E?p=29
*/

func main() {
	c := make(chan int, 3) // 带有缓冲的channel

	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("sub goroutine end")
		for i := 0; i < 4; i++ {
			c <- i // i写入channel
			fmt.Println("sub goroutine is running, len(c) = ", len(c), ", cap(c) = ", cap(c))
		}
	}()
    // sub goroutine 结束一定是在 num = 0 之后发生的
	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c // 从c中接收数据，并赋值为num
		fmt.Println("num = ", num)
	}

	fmt.Println("main end")
}
