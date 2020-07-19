package main

import "fmt"

/*

http://127.0.0.1:3999/methods/14
https://tour.go-zh.org/methods/14

空接口
指定了零个方法的接口值被称为  空接口。
如：
interface{}

空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）

空接口被用来处理未知类型的值。

例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。

*/

func main() {
	var i interface{}

	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = true
	describe(i)
}


func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}