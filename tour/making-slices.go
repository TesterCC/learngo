package main

import "fmt"

/*
http://127.0.0.1:3999/moretypes/13
https://tour.go-zh.org/moretypes/13

Creating a slice with make
用 make 创建切片
切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays

make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
a := make([]int, 5)  // len(a)=5

要指定它的容量，需向 make 传入第三个参数：
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
*/


func main() {
	// make 函数会分配一个元素为零值的数组并返回一个引用了它的切片
	a := make([]int, 5)
	printSlice2("a", a)

	// 要指定它的容量，需向 make 传入第三个参数
	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)

	/*
	a len=5 cap=5 [0 0 0 0 0]
	b len=0 cap=5 []
	c len=2 cap=5 [0 0]
	d len=3 cap=3 [0 0 0]
	*/
}


func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
