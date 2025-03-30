package main

import (
	"fmt"
	"math"
)

/*
https://tour.go-zh.org/methods/9
https://go.dev/tour/methods/9

接口

接口类型 的定义为一组方法签名。

接口类型的变量可以持有任何实现了这些方法的值。
*/

type Abser interface {
	Abs() float64
}

type MyFloat float64

// 实现了Abs()
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// 实现了Abs()
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// 接口类型的变量，可以持有任何实现了这些方法的值。
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat 实现了Abser
	//fmt.Printf("type of a = %T\n", a)
	a = &v // a *Vertex 实现了Abser
	//fmt.Printf("type of a = %T\n", a)

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v

	fmt.Println(a.Abs())
}
