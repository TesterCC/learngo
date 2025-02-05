package main

import "fmt"

// 注意类型在 变量名 之后。  https://blog.go-zh.org/gos-declaration-syntax
func add(x int, y int) int {
	return x + y
}

// 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
func add2(x, y int) int {
	return x + y
}


func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(11, 22))
}
