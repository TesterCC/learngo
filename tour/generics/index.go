package main

import "fmt"

/*

Go 在 1.18 之后才加入了泛型。

https://tour.go-zh.org/generics/1
https://go.dev/tour/generics/1


类型参数

可以使用 类型参数 编写Go函数 来处理多种类型。 函数的类型参数 出现在 函数参数之前的方括号之间。

func Index[T comparable](s []T, x T) int

此声明意味着 s 是满足内置约束 comparable 的任何类型 T 的切片。 x 也是相同类型的值。

comparable 是一个有用的约束，它能让我们对任意满足该类型的值使用 == 和 != 运算符。

在此示例中，我们使用它将值与所有切片元素进行比较，直到找到匹配项。 该 Index 函数适用于任何支持比较comparable的类型。

这段 Go 代码展示了 泛型（Generics） 的使用，主要用于编写支持多种数据类型的函数。
在 Python 中，你可能会使用 Duck Typing 或 typing.Generic 来实现类似的功能。

Python实现版本：~/Development/ws_python/Python3Scripts/dev_demo/generics_demo/generics_demo.py
*/

// Index 返回 x 在 s 中的下标，未找到则返回 -1。
// [T comparable] T 是泛型类型参数，表示 Index 可以适用于 任意支持 == 和 != 的类型（比如 int、string 等）。
// comparable 是 Go 预定义的约束，要求 T 必须支持比较（即 == 和 != 运算符）
// 参数 s[]T， 这是一个切片（slice），存储 T 类型的元素，类似于 Python 的 list[T]。
// x 是 T 类型的一个值，我们要在 s 中查找它。
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v 和 x 的类型为 T，它拥有 comparable 可比较的约束，
		// 因此我们可以使用 ==。
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index 可以在整数切片上使用
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15)) // 2

	// Index 也可以在字符串切片上使用
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello")) // -1
}
