package main

import (
	"fmt"
	"math"
)

/*
http://127.0.0.1:3999/methods/3
https://tour.go-zh.org/methods/3

You can declare a method on non-struct types, too.

在此例中，我们看到了一个带 Abs 方法的数值类型 MyFloat。你只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。
You can only declare a method with a receiver whose type is defined in the same package as the method.

（译注：就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。）
 */


type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
