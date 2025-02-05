package main

import "fmt"

/*
4种定义slice的方式
https://www.bilibili.com/video/BV1gf4y1r79E?p=13
*/
func main() {
	// 1.声明slice1是一个切片，初始化默认值未1，2，3，长度len是3
	//slice1 := []int{1, 2, 3}

	// 2.声明slice1是一个切片，但是并没有给slice分配空间
	//var slice1 []int
	////slice1[0] = 1  // report error，because slice1 no value
    //// 用make开辟内存空间，默认值0
	//slice1 = make([]int, 3)
	//slice1[0] = 100  // 这时赋值能成功

	// 3.声明slice1是一个切片，同时给slice分配空间，3个空间，初始化值为0
	//var slice1 []int = make([]int, 3)

	// 4.声明slice1是一个切片，同时给slice分配空间，3个空间，初始化值为0，通过:=推导出slice是一个切片
	slice1 := make([]int,3)
	// %v detail info
	fmt.Printf("len = %d, slice= %v\n", len(slice1), slice1)

	// 判断一个slice是否为0（空），是指是否有空间
	if slice1 == nil {
		fmt.Println("slice1 is empty slice")
	} else {
		fmt.Println("slice1 has space")
	}
}
