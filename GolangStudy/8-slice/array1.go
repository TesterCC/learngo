package main

import "fmt"

// 数组长度是固定的，值传递/拷贝传递
// 传递一个数组，实际建议传动态数组
func printArray(myArray [4]int) {
	// 值拷贝
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
	myArray[0] = 1111   //  并不会修改函数外的myArray3
}

func main() {
	// 固定长度饿数组
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index: ", index, ", value :", value)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray1 type = %T\n", myArray1)
	fmt.Printf("myArray2 type = %T\n", myArray2)
	fmt.Printf("myArray3 type = %T\n", myArray3)

	printArray(myArray3)
	fmt.Println("==========")
	for index, value := range myArray3 {
		fmt.Println("index: ", index, ", value :", value)
	}
}
