package main

import "fmt"

/*
28-Channel的基本定义与使用

channel用于goroutine之间通信

ref:
https://www.bilibili.com/video/BV1gf4y1r79E?p=28
*/
// main goroutine 和 sub goroutine 通过 c 进行交互
func main() {
	// 定义一个无缓冲channel
	c := make(chan int)

	go func() {
		defer fmt.Println("sub goroutine end")

		fmt.Println("goroutine 正在运行...")

		c <- 666 // 将666发送给c
	}()

	num := <-c // 从c中接受数据，并赋值为num  // c 使得 defer 永远在这句之后执行

	fmt.Println("num = ", num)
	fmt.Println("main goroutine end ...")
}
