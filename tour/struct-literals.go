package main

import "fmt"

/*
http://127.0.0.1:3999/moretypes/5
https://tour.go-zh.org/moretypes/5

结构体文法

结构体文法通过直接列出字段的值来新分配一个结构体。

使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）

特殊的前缀 & 返回一个指向结构体的指针。
*/

type Vertex3 struct {
	X, Y int
}


var (
	v1 = Vertex3{1,2}   // 创建一个 Vertex 类型的结构体
	v2 = Vertex3{X:1}    // Y:0 被隐式(implicit)地赋予
	v3 = Vertex3{}       // X:0 Y:0  直接隐式赋予
	p  = &Vertex3{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)
func main() {
	fmt.Println(v1, p, v2, v3)
}
