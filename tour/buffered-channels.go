package main

import "fmt"

/*
http://127.0.0.1:3999/concurrency/3
https://tour.go-zh.org/concurrency/3

Buffered Channels
Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道

ch := make(chan int, 100)

Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
*/

func main() {
	ch := make(chan int, 2)  // 缓冲长度不够会引起死锁，fatal error: all goroutines are asleep - deadlock!
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
