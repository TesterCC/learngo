package main

import "fmt"

/*
P50-51
数组，一段存储固定类型固定长度的连续内存空间。
数组大小不可变，数组大小必须指定，数组成员类型可以为任意类型。
数组成员可以修改。
*/

func main() {
	var classMates1 [3]string
	classMates1[0] = "Lily"      // 通过下表为数组成员赋值
	classMates1[1] = "Bob"
	classMates1[2] = "Tom"

	fmt.Println(classMates1)
	fmt.Println("The No.1 student is "+ classMates1[0])   // 通过下标访问数组成员

	// 使用初始化列表初始化数组
	classMates2 := [...]string{"Jack", "Alice", "Fred"}
	fmt.Println(classMates2)

	// 使用 指针 操作数组
	classMates3 := new([3]string)   // 通过new申请内存空间，并返回了其对应的指针

	classMates3[0] = "Alan"
	classMates3[1] = "Lisa"
	classMates3[2] = "Jerry"

	//fmt.Println(classMates3) //  &[Alan Lisa Jerry]
	fmt.Println(*classMates3)


}
