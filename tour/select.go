package main

import "fmt"

/*
http://127.0.0.1:3999/concurrency/5
https://tour.go-zh.org/concurrency/5

The select statement lets a goroutine wait on multiple communication operations.
select 语句使一个 goroutine 可以等待多个通信操作。

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。
*/

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)

}
