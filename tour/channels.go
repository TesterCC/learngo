package main

import "fmt"

/*
http://127.0.0.1:3999/concurrency/2
https://tour.go-zh.org/concurrency/2


信道    Channels

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。

ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.

和映射与切片一样，信道在使用前必须创建：
ch := make(chan int)

默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

以下示例对切片中的数进行求和，将任务分配给两个 Go Channel。一旦两个 Go Channel完成了它们的计算，它就能算出最终的结果。
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //  send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从c中接收

	fmt.Println(x, y, x+y)
}
