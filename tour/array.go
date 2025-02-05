package main

import "fmt"

/*
http://127.0.0.1:3999/moretypes/6
https://tour.go-zh.org/moretypes/6


数组
类型 [n]T 表示拥有 n 个 T 类型的值的数组。

表达式

var a [10]int
会将变量 a 声明为拥有 10 个整数的数组。

数组的长度是其类型的一部分，因此数组不能改变大小。
这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组。
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
