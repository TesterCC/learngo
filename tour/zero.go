package main

import "fmt"

/*

https://tour.go-zh.org/basics/12
https://go.dev/tour/basics/12

零值

没有明确初始化的变量声明会被赋予对应类型的 零值。

零值是：

数值类型为 0，
布尔类型为 false，
字符串为 ""（空字符串）

*/

func main() {
	var i int
	var f float64
	var f2 float32
	var b bool
	var s string
	fmt.Printf("%v %v %v %v %q\n", i, f, f2, b, s)
}
