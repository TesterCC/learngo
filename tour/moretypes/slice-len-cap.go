package main

import "fmt"

/*
https://tour.go-zh.org/moretypes/11
https://go.dev/tour/moretypes/11

切片的长度与容量:

切片拥有 长度 和 容量。

切片的长度就是它所包含的元素个数。

切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。

切割前面会影响cap即容量大小。

你可以通过重新切片来扩展一个切片，给它提供足够的容量。 试着修改示例程序中的切片操作，向外扩展它的长度，看看会发生什么。
*/

func main() {
	// normal
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // 丢弃前2个元素后，容量也缩小了

	// 注意slice s一直随重新切片而变化
	s = s[1:]
	printSlice(s) // len=1 cap=3 [7]

	fmt.Println("============================")
	t := []int{2, 4, 6, 8, 10}
	printSlice(t) // len=5 cap=5 [2 4 6 8 10]

	t = t[2:]
	printSlice(t) // len=3 cap=3 [6 8 10]

	t = t[:1]
	printSlice(t) // len=1 cap=3 [6]

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
