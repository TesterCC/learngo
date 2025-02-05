package main

import "fmt"

/*
http://127.0.0.1:3999/moretypes/2
https://tour.go-zh.org/moretypes/2

http://127.0.0.1:3999/moretypes/3
https://tour.go-zh.org/moretypes/3

结构体
一个结构体（struct）就是字段（field）的集合。
A struct is a collection of fields.

结构体字段 Struct Fields
结构体字段使用点号来访问。
*/

// 定义结构体
type Vertex struct {
	X int
	Y int
}


// 结构体字段
func main() {
	fmt.Println(Vertex{})
	fmt.Println(Vertex{1,2})
	fmt.Println("--------------------")

	v := Vertex{3,4}
	fmt.Println(v)
	v.X = 7
	fmt.Println(v.X)
}
