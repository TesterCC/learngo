package main

import (
	"fmt"
	"sync"
)

/*
https://www.bilibili.com/video/BV1SjX4YhEas/

变形题：
https://www.nowcoder.com/discuss/712384113688199168

go协程奇偶数交替打印

使用2个协程交替打印 1和N 之间的奇偶数，保证输出顺序严格按照 奇数 -> 偶数 -> 奇数 -> 偶数  的顺序交替进行
*/

// 打印奇数
func printOdd(max int, wg *sync.WaitGroup, chan1, chan2 chan struct{}) {
	defer wg.Done() // 在协程完成时将其关闭
	for i := 1; i <= max; i += 2 {
		// chan接收到信号才打印，不然不打印。
		<-chan1
		fmt.Println("odd: ", i)
		chan2 <- struct{}{}
	}
}

// 打印偶数
func printEven(max int, wg *sync.WaitGroup, chan1, chan2 chan struct{}) {
	defer wg.Done()
	for i := 2; i <= max; i += 2 {
		<-chan2
		fmt.Println("even: ", i)
		chan1 <- struct{}{}
	}
}

func main() {
	chan1 := make(chan struct{}, 1)
	chan2 := make(chan struct{}, 1)

	var wg sync.WaitGroup
	wg.Add(2)
	// 利用管道在协程中进行通信
	// 因为是从奇数开始，所以启动信号发送到chan1
	chan1 <- struct{}{}
	go printOdd(100, &wg, chan1, chan2)
	go printEven(100, &wg, chan1, chan2)
	wg.Wait()
}
