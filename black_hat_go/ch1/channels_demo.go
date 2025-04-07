package main

import "fmt"

// P16-17 use channels
// Go 通道（channel）数据类型提供了一种机制，通过该机制，goroutine可以同步执行函数并且这些函数可以相互通信。

// 接收2个参数，即一个字符串，一个用于同步数据的通道
func strlen(s string, c chan int) {
	// 将len()结果放入通道，即流入通道
	c <- len(s)
}

// demo：使用通道同时显示不同字符串的长度及其总和。
func main() {
	// 可以定义各种类型的通道，具体取决于要在通道传递的数据类型。
	// 定义并使用chan int类型的变量c，传递整数
	c := make(chan int)

	// 进行多个并发调用，这将启动多个goroutine。
	go strlen("Salutations", c) // 11
	go strlen("World", c)       // 5
	go strlen("Test", c)        // 4

	// <- 指示数据是流入还是流出通道，这里是流出通道
	// 从通道读取数据
	x, y, z := <-c, <-c, <-c
	fmt.Println(x, y, z, x+y+z)
}
