package main

import "fmt"

/*
https://tour.go-zh.org/moretypes/2
https://go.dev/tour/moretypes/2

https://tour.go-zh.org/moretypes/3
https://go.dev/tour/moretypes/3

结构体
一个 结构体（struct）就是一组 字段（field）。

一个结构体（struct）就是字段（field）的集合。
A struct is a collection of fields.

结构体字段 Struct Fields

结构体字段可通过点号 . 来访问。

structs.go    structs-fields.go
*/

// 定义结构体
type Vertex struct {
	X int
	Y int
}

type TreeD struct {
	X, Y, Z float64
	Name    string
}

// 结构体字段
func main() {
	fmt.Println(Vertex{})     // {0 0}
	fmt.Println(Vertex{1, 2}) // {1 2}
	fmt.Println("--------------------")

	v := Vertex{3, 4}
	fmt.Println(v)
	v.X = 7
	fmt.Println(v.X)

	fmt.Println("--------custom test------------")
	// custom struct
	td := TreeD{2, 3.3, 6.6, "Tester"} // {2 3.3 6.6 Tester}

	fmt.Println(td)
	fmt.Println(td.Name)
	fmt.Println(td.Z) // 6.6
}
