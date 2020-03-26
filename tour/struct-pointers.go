package main

import "fmt"

/*
https://tour.go-zh.org/moretypes/4
http://127.0.0.1:3999/moretypes/4

Pointers to structs 结构体指针

结构体字段可以通过结构体指针来访问。

如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
*/

type Vertex2 struct {
	X int
	Y int
}

func main() {

	v := Vertex2{1, 2}
	p := &v   // & 操作符会生成一个指向v操作数的指针
	p.X = 1e9
	p.Y = 2<<9 // 左移表示乘以2的几次方
	//p.Y = 3<<3  // 左移，这里表示表示3乘以2的3次方 3 * 2^3 = 3 * 8 = 24
	fmt.Println(v)

}
