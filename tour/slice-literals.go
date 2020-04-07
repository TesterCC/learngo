package main

import "fmt"

/*
http://127.0.0.1:3999/moretypes/9
https://tour.go-zh.org/moretypes/9

Slice literals
切片文法
切片文法类似于没有长度的数组文法。

一个数组文法:
[3]bool{true, true, false}

创建一个和上面相同的数组，然后构建一个引用了它的切片:
[]bool{true, true, false}
*/

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

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
}
