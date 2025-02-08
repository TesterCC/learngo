package main

import "fmt"

/*
https://tour.go-zh.org/basics/8
https://go.dev/tour/basics/8
var 语句用于声明一系列变量。和函数的参数列表一样，类型在最后。var 语句可以出现在包或函数的层级。
*/
var c, python, java, rust bool // a list of variables

func main() {
	var i int
	fmt.Println(i, c, python, java, rust)
	// default init: 0 false false false false
}
