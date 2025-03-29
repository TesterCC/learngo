package main

import "fmt"

/*
https://tour.go-zh.org/moretypes/19
https://go.dev/tour/moretypes/19

map 映射

A map maps keys to values.

map 映射将键映射到值。

映射的零值为 nil 。nil 映射既没有键，也不能添加键。

The zero value of a map is nil. A nil map has no keys, nor can keys be added.
make 函数会返回给定类型的映射map，并将其初始化备用。
*/

type Vertex struct {
	Lat, Long float64 // Geo Location (纬度, 经度)
}

var m map[string]Vertex // 声明一个 map，键为 string，值为 Vertex 结构体

func main() {
	fmt.Println(m)              // 1. 打印 map 的初始值
	m = make(map[string]Vertex) // 2. 使用 make() 初始化 map
	m["Pentest Lab"] = Vertex{  // 3. 插入键值对
		40.68433, -74.39967,
	}
	fmt.Println(m)                // 4. 打印 map 的内容   // map[Pentest Lab:{40.68433 -74.39967}]
	fmt.Println(m["Pentest Lab"]) // 5. 访问并打印 "Pentest Lab" 对应的值 // {40.68433 -74.39967}
}
