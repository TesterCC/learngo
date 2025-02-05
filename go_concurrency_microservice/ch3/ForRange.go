package main

import "fmt"

/*
P57
对 数组、切片、字典进行遍历
*/

func main() {
	// 数组的遍历
	nums := [...]int{1,2,3,4,5,6,7,8}
	for k,v := range nums {
		// k为下标，v为对应的值
		fmt.Println(k,v," ")
	}
	fmt.Println()

	// 切片的遍历
	slis := []int{1,2,3,4,5,6,7,8}
	for k,v := range slis {
		// k为下标，v为对应的值
		fmt.Println(k,v," ")
	}

	fmt.Println()

	// 字典的遍历
	tmpMap := map[int]string{
		0: "Alice",
		1: "Bob",
		2: "Chris",
	}

	for k,v := range tmpMap{
		//  k为键的值，v为对应值
		fmt.Println(k,v," ")
	}
}
