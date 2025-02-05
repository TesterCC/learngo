package main

import "fmt"

var length float64 = 1.2 // int 2
var width int = 2

var length2 float64 = 3.75
var width2 int = 5

func main() {
	// int to float
	fmt.Println("int to float: ", float64(width))
	fmt.Println("Area is ", length*float64(width))
	fmt.Println("length > width ? ", length > float64((width))) // 比较对应要为同类型，否则报错。

	// float to int
	width2 = int(length2) // 该转换将导致小数部分被删除
	fmt.Println("float to in: ", width2)
}
