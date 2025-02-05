package main

import (
	"fmt"
	"math"
)

/*

http://127.0.0.1:3999/methods/11
https://tour.go-zh.org/methods/11

接口也是值。它们可以像其它值一样传递。

接口值 可以用作函数的 参数或返回值。

在内部，接口值可以看做包含值和具体类型的元组：

(value, type)

接口值保存了一个具体底层类型的具体值。

接口值 调用方法 时会执行其 底层类型的同名方法。

*/

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

}
