package main

import "fmt"

// 3-1 数组 array, 数组是值类型

// 注意区别，[5]int是一个数组，[]int是切片
func printArray(arr [5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

func printArray2(arr *[5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	var arr1 [5]int //指定长度不指定默认值初始化

	arr2 := [3]int{1, 3, 5} //指定长度和内容初始化

	arr3 := [...]int{2, 4, 6, 8, 10} //不指定长度初始化

	// 定义二维数组，数量写在类型前面   grid网格
	var grid [4][5]int // 4行5列，从前往后读
	var grid2 [4][5]bool

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	fmt.Println(grid2)

	// 遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	// 用range关键字遍历数组   i获得数组下标
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	// 用range关键字遍历数组   i获得数组下标，v获得下标对应的值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// 用range关键字遍历数组  只要数值v
	for _, v := range arr3 {
		fmt.Println(v)
	}

	fmt.Println("Print arr1:")
	printArray(arr1)
	printArray2(&arr1)
	fmt.Println("Print arr2:")
	printArray(arr3)
	printArray2(&arr3)
	fmt.Println("==================================")
	fmt.Println(arr1, arr3)   // a[0]=100 在printArray受影响，在外部不受影响。 如果是传递的指针, a[0]=100 会受影响
}
