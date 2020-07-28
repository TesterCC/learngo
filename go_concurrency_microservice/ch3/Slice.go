package main

import "fmt"

/*
P51-54

切片Slice是对数组一个连续片段的引用，Slice是一个容量可变的序列

可以简单将切片理解为动态数组

Slice内部结构：底层数组指针array、切片长度len、切片容量cap。   cap >= len

*/

func main() {
	// 1.从原生数组中生产一个切片，那么 生成的切片指针 即指向 原数组
	source := [...]int{1, 2, 3}
	sli := source[0:1] // 切片指针 指向 原数组
    fmt.Println("1.从原生数组中生产一个切片")
	fmt.Printf("slice value is %v\n", sli)
	fmt.Printf("slice len is %v\n", len(sli))
	fmt.Printf("slice cap is %v\n", cap(sli))

	// 此时，因为Slice作为指向原有数组的引用，对切片内成员进行修改，就是对原有数组进行修改
	sli[0] = 4
	fmt.Printf("slice value is %v\n", sli)
	fmt.Printf("source value is %v\n", source)

	// 2.动态创建切片
	// 通过make函数动态创建切片，在创建过程中指定切片的长度和容量    make([]T,size,cap)
	slic := make([]int,2,4)   // 成员类型、长度、容量
	fmt.Println("2.动态创建切片")
	fmt.Printf("slice value is %v\n", slic)    // 切片成员被初始化为指定类型的初始值
	fmt.Printf("slice len is %v\n", len(slic))
	fmt.Printf("slice cap is %v\n", cap(slic))

	// 3.声明新的切片  var name []T
	ex := []int{1,2,3}   // []不需要指定其大小
	fmt.Println("3.声明新的切片")
	fmt.Printf("ex value is %v\n", ex)
	fmt.Printf("ex len is %v\n", len(ex))
	fmt.Printf("ex cap is %v\n", cap(ex))
}
