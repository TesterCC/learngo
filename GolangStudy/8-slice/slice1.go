package main

import "fmt"

/*
slice切片，即动态数组
1.golang默认都是采用值传递，即拷贝传递
2.有些值天生就是指针（slice、map、channel）
*/
func printArray2(myArray []int) {
	// 引用传递
	// _  表示匿名变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 999
}

func main() {
	myArray := []int{1, 2, 3, 4} // 动态数组，切片slice 未限制数组长度
	fmt.Printf("myArray type is %T\n", myArray)
	printArray2(myArray)

	fmt.Println("==========")
	//printArray2(myArray)
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
