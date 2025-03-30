package main

import (
	"fmt"
	"math"
)

/*
https://tour.go-zh.org/methods/1
https://go.dev/tour/methods/1

方法

Go 没有类(Class)。不过你可以为类型(Type)定义方法。

方法就是一类带特殊的 接收者 参数的函数。
A method is a function with a special receiver argument.

方法接收者在它自己的参数列表内，位于 func关键字 和 方法名 之间。

在此例中，Abs 方法拥有一个名字为 v，类型为 Vertex0 的接收者。
*/

type Vertex0 struct {
	X, Y float64
}

/*
方法Method就是一类带特殊的 接收者receiver 参数的函数。
方法和普通函数有区别：
在 Go 语言中，方法（Method）与普通函数的区别在于方法必须有一个特殊的参数——接收者（Receiver）。
	- 普通函数： 没有 receiver，直接传递参数调用。
	- 方法（Method）： 绑定在一个特定的类型（struct、interface 等）上，并且有一个 receiver。
*/

// 接收者receiver在此为 v Vertex0
func (v Vertex0) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex0{3, 4}
	fmt.Println(v.Abs())
}
