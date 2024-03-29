package main

import (
	"fmt"
	"math"
)

/*
http://127.0.0.1:3999/methods/2
https://tour.go-zh.org/methods/2

方法即函数
记住：方法只是个带接收者参数的函数。
Remember: a method is just a function with a receiver argument.

现在这个 Abs 的写法就是个正常的函数，功能并没有什么变化。
Here's Abs written as a regular function with no change in functionality.

*/

type Vertex struct {
	X,Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3,4}
	fmt.Println(Abs(v))

}
