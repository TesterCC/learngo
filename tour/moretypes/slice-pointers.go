package main

import "fmt"

/*
https://tour.go-zh.org/moretypes/8
https://go.dev/tour/moretypes/8

切片Slice类似数组Array的引用

切片就像数组的引用 切片并不存储任何数据，它只是描述了底层数组中的一段。

更改切片的元素会修改其底层数组中对应的元素。

和它共享底层数组的切片都会观测到这些修改。
*/

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println("origin array: ", names)

	// slice
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	fmt.Println("slice a: ", a)
	fmt.Println("slice b: ", b)

	// b[0] 和 a[1] 共享底层
	b[0] = "XXX"
	fmt.Println(a, b)

	fmt.Println("slice a: ", a)
	fmt.Println("slice b: ", b)

	fmt.Println("current array: ", names)

}
