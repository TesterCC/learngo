package main

import "fmt"

/*
32-channel与select
ref:
https://www.bilibili.com/video/BV1gf4y1r79E?p=32

单流程下一个go只能监控一个channel的状态，select可以完成监控多个channel的状态
select具备多路channel的监控状态功能
*/

// 计算斐波那契数列
func fibonacii(c,quit chan int)  {
	//channel传形参是传引用
	x,y := 1,1

	// 这里相当于while(true)
	for {
		select {
		case c <-x :
			// 若c可写，则该case就会进来
			x,y = y, x+y
		case <-quit:  //  当quit可读，说明0已写入quit
			fmt.Println("quit")
			return   // exit loop
		}
	}

}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // 从c中读数据并打印，若c没有数据则阻塞
		}

		quit <- 0 // 循环完成后，向quit中发送数据0，表示可退出
	}()

	// main go
	fibonacii(c,quit)
}
