package main

import "fmt"

/*
https://tour.go-zh.org/basics/6

多返回值
函数可以返回任意数量的返回值。

swap 函数返回了3个字符串。
*/

func swap(x, y, z string) (string, string, string) {
	return z, y, x
}

func main() {
	a, b, c := swap("hello", "world", "go")
	fmt.Println(a, b, c)
}
