package main

import "fmt"

/*
https://tour.go-zh.org/moretypes/9
https://go.dev/tour/moretypes/9

Slice literals

切片字面量

切片字面量类似于没有长度的数组字面量。

这是一个数组字面量：
[3]bool{true, true, false}

p.s. 数组字面量，注意，这里没有指定长度

下面这样则会创建一个和上面相同的数组，然后再构建一个引用了它的切片：
[]bool{true, true, false}

*/

func main() {

	//  slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// 匿名结构体切片 的用法
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// custom 匿名结构体切片
	m := []struct {
		name   string
		age    int
		online bool
	}{
		{"Alice", 18, true},
		{"Bob", 22, false},
		{"Chris", 33, true},
		{"David", 46, false},
	}
	fmt.Println(m)
}
