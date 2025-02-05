package main

import "fmt"

/*

https://tour.go-zh.org/basics/5

函数（续）
当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。

在本例中，

x int, y int
被简写为
x, y int

*/

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(222, 666))
}
