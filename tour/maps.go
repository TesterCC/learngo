package main

import "fmt"

/*
http://127.0.0.1:3999/moretypes/19
https://tour.go-zh.org/moretypes/19

A map maps keys to values.
The zero value of a map is nil. A nil map has no keys, nor can keys be added.
make 函数会返回给定类型的映射map，并将其初始化备用。
*/

type Vertex struct {
	Lat, Long float64 // Geo Location
}

var m map[string]Vertex

func main() {
	fmt.Println(m)
	m = make(map[string]Vertex)
	m["Pentest Lab"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Pentest Lab"])
}
