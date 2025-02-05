package main

import "fmt"

/*
在分配内存空间之前需要先确定新的切片容量，Go 语言根据切片的当前容量选择不同的策略进行扩容：

1.如果期望容量大于当前容量的两倍就会使用期望容量；
2.如果当前切片容量小于 1024 就会将容量翻倍；
3.如果当前切片容量大于 1024 就会每次增加 25% 的容量，直到新容量大于期望容量；
*/

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Printf("cap a: %d, len a: %d, a value: %v\n", cap(a), len(a), a)
	a = append(a, 11)
	fmt.Printf("cap a: %d, len a: %d, a value: %v\n", cap(a), len(a), a) // 扩容

	a1 := a
	a2 := a
	a1 = append(a1, 12)
	fmt.Printf("cap a1: %d, len a1: %d, a1 value: %v\n", cap(a1), len(a1), a1)
	fmt.Printf("cap a2: %d, len a2: %d, a2 value: %v\n", cap(a2), len(a2), a2)
	fmt.Println(a1, a2) // [1 2 3 4 11 12] [1 2 3 4 11]

	a1[0] = 111111
	fmt.Println(a1, a2) // [111111 2 3 4 11 12] [111111 2 3 4 11]
}
