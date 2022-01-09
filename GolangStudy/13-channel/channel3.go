package main

import "fmt"

/*
30-channel的关闭特点
ref:
https://www.bilibili.com/video/BV1gf4y1r79E?p=30
*/

// 2个goroutine通信，尝试关闭channel
func main() {
	// 创建无缓冲channel
	c := make(chan int)

	go func() {
		for i :=0; i<5; i++{
			c <- i  // 将i写入channel
		}

		// close可以关闭一个channel   这里 sub goroutine遍历完才会关闭channel
		close(c)
	}()

	// main goroutine
	for {
		// ok若为true表示channel没有关闭，若为false表示channel已经关闭
		if data,ok := <-c; ok {
			fmt.Println(data)
		} else {
			break  // 跳出当前的死循环
		}
	}

	fmt.Println("Main Finished ...")
}