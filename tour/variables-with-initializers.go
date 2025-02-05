package main

import "fmt"

/*
https://tour.go-zh.org/basics/9
https://go.dev/tour/basics/9
变量的初始化
变量声明可以包含初始值，每个变量对应一个。、
如果提供了初始值，则类型可以省略；变量会从初始值中推断出类型。
*/

// var i,j int = 1,2 // If an initializer is present, the type can be omitted
var i, j = 1, 2

func main() {
	var c, python, java, rust = true, false, "No!", 1
	fmt.Println(i, j, c, python, java, rust)
	// 1 2 true false No! 1
}
