package main

import "fmt"

/*
https://tour.go-zh.org/moretypes/25
https://go.dev/tour/moretypes/25

函数的闭包
Go函数可以是一个闭包。
闭包是一个函数值，它引用了其函数体之外的变量。
该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。

例如，函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。

闭包（Closure）
这段 Go 代码展示了闭包的概念。闭包是一个函数，它可以访问外部作用域的变量，并在多次调用时保持变量的状态。
	- adder() 返回一个匿名函数（闭包），这个函数接受 int 类型参数 x，并将其累加到 sum 上。
	- sum 变量不会因 adder() 执行完毕而销毁，它会随着闭包一起存活，并在每次调用时继续累加。
这样，Go 通过闭包提供了一种持久化局部变量的能力，非常适合累加器、缓存、事件处理等应用。
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
