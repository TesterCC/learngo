package main

import "fmt"

/*

在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出。

REF: https://www.cnblogs.com/nulige/p/10201488.html   (go语言格式化输出)
*/


func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)   // %T使用Go语法输出的值的类型

	v1 := "test"
	fmt.Printf("v is of type %T\n", v1)

	v2 := false
	fmt.Printf("v is of type %T\n", v2)

	v3 := 3.1415926
	fmt.Printf("v is of type %T\n", v3)

	v4 := 0.666 + 0.7i
	fmt.Printf("v is of type %T\n", v4)
}

