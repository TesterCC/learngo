package main

import (
	"fmt"
	"math"
)

/*
http://127.0.0.1:3999/methods/1
https://tour.go-zh.org/methods/1

方法
Go 没有类。不过你可以为结构体类型定义方法。

方法就是一类带特殊的 接收者receiver 参数的函数。

方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。

在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。
*/

type Vertex struct {
	X,Y float64
}

// 方法就是一类带特殊的 接收者receiver 参数的函数。和普通函数有区别
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := Vertex{3,4}
	fmt.Println(v.Abs())
}
