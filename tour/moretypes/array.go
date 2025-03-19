package main

import "fmt"

/*
https://tour.go-zh.org/moretypes/6
https://go.dev/tour/moretypes/6


数组

类型 [n]T 表示一个数组，它拥有 n 个类型为 T 的值。

表达式

var a [10]int

会将变量 a 声明为拥有 10 个整数的数组。

数组的长度是其类型的一部分，因此 数组不能改变大小。

这看起来是个限制，不过没关系，Go拥有更加方便的使用数组的方式。
*/

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

}
