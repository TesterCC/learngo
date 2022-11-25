package main

import "fmt"

// P45-47
/*
Go指针包含3个概念：
- 指针地址
- 指针类型
- 指针取值
*/
func main() {
	str := "Golang is Good!"

	// get str pointer
	strPtr := &str

	fmt.Printf("str type is %T, and value is %v\n", str, str)
	fmt.Printf("strPtr type is %T, and value is %v\n", strPtr, strPtr)

	// 可以继续对指针进行取址操作
	strPtrPtr := &strPtr
	fmt.Printf("strPtrPtr type is %T, and value is %v\n", strPtrPtr, strPtrPtr)

	// 获取指针对应变量的值     & 取址    * 取值,指针
	// 赋值过程中发生了值拷贝，值拷贝会创建新的内存空间，然后将原有变量的值复制到新的内存空间中，形成两个独立的变量。
	newStr := *strPtr
	fmt.Printf("newStr type is %T, and value is %v, address is %p\n", newStr, newStr, &newStr)

	*strPtr = "Python is Good, too!"   // 通过指针对变量进行赋值

	fmt.Printf("newStr type is %T, and value is %v, address is %p\n", newStr, newStr, &newStr)

	fmt.Printf("str type is %T, and value is %v, address is %p\n", str, str, &str)

    str2 := new(string)

    *str2 = "PHP is Good!"
	fmt.Printf("str2 type is %T, and value is %v, address is %p\n", str2, str2, &str2)
}
