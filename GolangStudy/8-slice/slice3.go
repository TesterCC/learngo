package main

import "fmt"

/*
14 - slice切片追加与截取
ref: https://www.bilibili.com/video/BV1gf4y1r79E?p=14

调用append后需要扩容，会重新开辟一块内存分配给新的切片 而不是在原来的基础上扩容

ref: https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-array-and-slice/#324-%E8%BF%BD%E5%8A%A0%E5%92%8C%E6%89%A9%E5%AE%B9
在分配内存空间之前需要先确定新的切片容量，运行时根据切片的当前容量选择不同的策略进行扩容：
1.如果期望容量大于当前容量的两倍就会使用期望容量；
2.如果当前切片的长度小于 1024 就会将容量翻倍；
3.如果当前切片的长度大于 1024 就会每次增加 25% 的容量，直到新容量大于期望容量；
*/
func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	// 向numbers切片追加一个元素1
	numbers = append(numbers, 1)  // numbers len 4, [0 0 0 1], cap 5
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 2)  // numbers len 5, [0 0 0 1 2], cap 5
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向一个容量cap已经满的slice追加元素
	numbers = append(numbers, 3)  // numbers len 6, [0 0 0 1 2 3], cap 10
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("---------------------------")

	var numbers2 = make([]int,3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)

}
