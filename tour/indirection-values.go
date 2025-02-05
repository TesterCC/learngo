package main

/*

http://127.0.0.1:3999/methods/7
https://tour.go-zh.org/methods/7

同样的事情也发生在相反的方向。

接受一个值作为参数的函数必须接受一个指定类型的值：
而以值为接收者的方法被调用时，接收者既能为值又能为指针：
这种情况下，方法调用 p.Abs() 会被解释为 (*p).Abs()。
*/

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())   // 方法调用 p.Abs() 会被解释为 (*p).Abs()。
	fmt.Println(AbsFunc(*p))
}
