package main

import "fmt"

/*
https://tour.go-zh.org/moretypes/26
https://go.dev/tour/moretypes/26

练习：斐波纳契闭包
让我们用函数做些好玩的事情。

实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`。

在这个数列中，从第三项开始，每一项都等于前两项之和。
如果用F(n)表示第n项的值，那么可以用递推公式表示为F(n)=F(n - 1)+F(n - 2)，其中F(0)=0，F(1)=1。
*/

// 返回一个“返回int的函数”
func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		tmp := a
		a, b = b, a+b
		return tmp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
