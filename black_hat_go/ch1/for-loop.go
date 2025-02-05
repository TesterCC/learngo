package main

import "fmt"

//P15

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 初始化一个整数切片 nums
	nums := []int{2, 4, 6, 8}
	for idx, val := range nums {
		fmt.Println(idx, val)
	}
}
