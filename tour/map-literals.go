package main

import "fmt"

/*
http://127.0.0.1:3999/moretypes/20
https://tour.go-zh.org/moretypes/20

Map literals are like struct literals, but the keys are required.


http://127.0.0.1:3999/moretypes/21
https://tour.go-zh.org/moretypes/21

映射的文法（续）
若顶级类型只是一个类型名，你可以在文法的元素中省略它。(看例子)

 */

type Vertex struct {
	Lat, Long float64
}

// define directly
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// 若顶级类型只是一个类型名，你可以在文法的元素中省略它。
var m2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
	fmt.Println(m2)
}
