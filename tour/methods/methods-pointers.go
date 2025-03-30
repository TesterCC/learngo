package main

import (
	"fmt"
	"math"
)

/*
https://tour.go-zh.org/methods/4
https://go.dev/tour/methods/4

指针类型的接收者

你可以为指针类型的接收者声明方法。
You can declare methods with pointer receivers.

这意味着对于某类型 T，接收者的类型可以用 *T 的文法。（此外，T 不能是像 *int 这样的指针。）

例如，这里为 *Vertex3 定义了 Scale 方法。

指针接收者的方法可以修改接收者指向的值（就像 Scale 在这做的）。由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。

*/

type Vertex3 struct {
	X, Y float64
}

// Value Receiver
// Abs() 方法的接收者是 Vertex3，它不会修改原对象，只是计算 X 和 Y 的欧几里得范数（模）。
func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Pointer Receiver
// Scale(f float64) 方法的接收者是 *Vertex3，意味着它可以修改原对象;直接对 v.X 和 v.Y 进行缩放计算，影响原 Vertex3 实例。
func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex3{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

}
