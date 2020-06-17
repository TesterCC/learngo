package main

import "fmt"

/*
3-2 切片的概念
https://www.bilibili.com/video/BV18Q4y1M7NV?p=10
todo
 */

func updateSlice(s []int){
	s[0] = 100
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	arr2 := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr2[2:]
	fmt.Println("new arr[2:] = ", s1)
	s2 := arr2[:]
	fmt.Println("new arr[:] = ", s2)
}
