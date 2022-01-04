package main

import "fmt"

/*

ref: https://www.bilibili.com/video/BV1gf4y1r79E?p=11

*/

func main() {
	// 写入defer关键字  在执行最后，在结束之前调用
	// 多个defer是压栈的形式，后进先出
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")

}

/*
main::hello go 1
main::hello go 2
main end2
main end1
*/
