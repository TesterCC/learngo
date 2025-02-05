package main

import (
	"fmt"
	"time"
)

// P92-93 4-4 使用switch从多个channel中读取数据
func send(ch chan int, begin int)  {
	// 循环向通道发送数据
	for i := begin; i < begin + 10; i++ {
		ch <- i
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go send(ch1, 0)
	go send(ch2, 10)

	// 主 goroutine 休眠1s，保证调度成功
	time.Sleep(time.Second)

	for {
		// 通过select多路复用分别从ch1和ch2中读取数据，如果多个case语句中的ch同时到达，那select将会运行一个伪随机算法选择一个case。
		select {
		case val := <- ch1:   // 从 ch1 读取数据
			fmt.Printf("get value %d from ch1\n", val)
		case val := <- ch2:  // 从 ch2 读取数据
			fmt.Printf("get value %d from ch2\n", val)
		case <- time.After(2 * time.Second):   // 超时设置， 由于channel的阻塞时无法被中断的，所以这是一种有效地从阻塞的channel中超时返回的小技巧。
			fmt.Println("Time out")
			return
		}
	}
}

// 由于goroutine调度的不确定性，输出结果顺序可能变化。