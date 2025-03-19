package main

import "fmt"

/*
https://tour.go-zh.org/moretypes/15
https://go.dev/tour/moretypes/15


向切片追加元素
为切片追加新的元素是种常见的操作，为此 Go 提供了内置的 append 函数。内置函数的文档对该函数有详细的介绍。


func append(s []T, vs ...T) []T
append 的第一个参数 s 是一个元素类型为 T 的切片，其余类型为 T 的值将会追加到该切片的末尾。

append 的结果是一个包含原切片所有元素加上新添加元素的切片。

当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。

ref:
https://go-zh.org/pkg/builtin/#append

en：https://go.dev/blog/slices-intro
https://blog.go-zh.org/go-slices-usage-and-internals

*/

func main() {
	var s []int
	printSlice3(s) // len=0 cap=0 []

	// add a empty slice
	s = append(s, 0)
	printSlice3(s) // len=1 cap=1 [0]

	// 这个切片会按需增长
	s = append(s, 1) // len=2 cap=2 [0 1]
	printSlice3(s)

	// 可以一次性添加多个元素
	s = append(s, 22, 33, 44, 55) // len=6 cap=6 [0 1 22 33 44 55]
	printSlice3(s)
}

func printSlice3(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
