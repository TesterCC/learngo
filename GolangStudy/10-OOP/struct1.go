package main

import "fmt"

/*
17 - struct基本定义与使用
ref: https://www.bilibili.com/video/BV1gf4y1r79E?p=17
*/

// 声明一种行的数据类型 myint， 是int的一个别名
type myint int

func main() {
	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)   //  main.myint
}
