package main

import (
	"fmt"
	"time"
)

//P91-92  4-3 使用带缓冲区的channel

func consume(ch chan int){
	// 线程休息 7s 再， 从 channel 读取数据
	time.Sleep(time.Second * 7)
	<- ch  // 从ch读取数据
}

func main() {
	// 创建一个长度为2的channel
	ch := make(chan int, 2)
	go consume(ch)

	ch <- 0
	ch <- 1
	// 发送数据不被阻塞
	fmt.Println("I am free!")
	ch <- 2
	fmt.Println("I can not go there within 7s!")

	time.Sleep(time.Second)
}
