package main

import (
	"fmt"
	"math"
)

/*
https://tour.go-zh.org/methods/2
https://go.dev/tour/methods/2

方法即函数
记住：方法method只是个带接收者receiver参数的函数。
Remember: a method is just a function with a receiver argument.

现在这个 Abs 的写法就是个正常的函数，功能并没有什么变化。
Here's Abs written as a regular function with no change in functionality.

*/

type Vertex2 struct {
	X, Y float64
}

func Abs2(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex2{3, 4}
	fmt.Println(Abs2(v))

}
