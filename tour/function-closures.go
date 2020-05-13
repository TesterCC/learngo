package main

import "fmt"

/*
http://127.0.0.1:3999/moretypes/25
https://tour.go-zh.org/moretypes/25

函数的闭包
Go 函数可以是一个闭包。
闭包是一个函数值，它引用了其函数体之外的变量。
该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。

例如，函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。
*/

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
