package main

import (
	"fmt"
	"math"
)

/*

https://tour.go-zh.org/methods/11
https://go.dev/tour/methods/11

接口值

接口也是值。它们可以像其它值一样传递。

接口值 可以用作函数的参数或返回值。

在内部，接口值可以看做包含值和具体类型的元组：
(value, type)

接口值保存了一个具体底层类型的具体值。

接口值调用方法时会执行其底层类型的同名方法。

这段 Go 代码演示了 接口（interface） 作为值的特性，以及如何使用接口来存储不同的类型，同时调用它们的方法。
*/

type I interface {
	M()
}

type T struct {
	S string
}

// 任何类型只要实现了 M() 方法，就被认为实现了 I 接口。
// func (t *T) M() 使用指针接收者，实现了 I 接口；意味着只有 *T（T 的指针）才能被赋值给 I 类型，而 T 本身不行。
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

// F 是一个自定义类型（float64 的别名）。
// func (f F) M() 使用的是值接收者：
// F 直接实现了 M() 方法，所以值类型和指针类型都可以赋给 I。
func (f F) M() {
	fmt.Println(f)
}

// describe(i I) 接收一个 I 接口，并打印它的具体值和具体类型。
func describe(i I) {
	// %v 输出值，%T 输出类型。
	fmt.Printf("(%v,%T)\n", i, i)
}

func main() {
	var i I // 声明接口变量 i

	i = &T{"Hello"} // 赋值：*T 类型实现了 I
	describe(i)     // 输出值和类型
	i.M()           // 调用 T 的 M() 方法，打印 "Hello"

	fmt.Println("----------------")

	i = F(math.Pi) // 赋值：F 类型实现了 I
	describe(i)    // 输出值和类型
	i.M()          // 调用 F 的 M() 方法，打印 3.141592653589793

}
