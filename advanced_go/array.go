package main

import "fmt"

// P22 数组的定义方式

func main() {
	var a [3]int   // [0 0 0]
	var b = [...]int{1,2,3}   // Go 语言会在编译期间通过源代码对数组的大小进行推断
	var c = [...]int{2:3,1:2}  // [0 2 3]    // 数组长度以出现的最大索引为准，没有明确初始化的与纳素依然用0    map[int]Type
	var d = [...]int{1,2,4:5,6} // [1 2 0 0 5 6]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
