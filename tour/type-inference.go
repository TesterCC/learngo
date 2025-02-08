package main

import "fmt"

/*
https://tour.go-zh.org/basics/14
https://go.dev/tour/basics/14

类型推断

在声明一个变量而不指定其类型时（即使用不带类型的 := 语法 var = 表达式语法），变量的类型会通过右值推断出来。

当声明的右值确定了类型时，新变量的类型与其相同：

var i int
j := i // j 也是一个 int
不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 int、float64 或 complex128 了，这取决于常量的精度：

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
试着修改示例代码中 v 的初始值，并观察它是如何影响类型的。

REF: https://www.cnblogs.com/nulige/p/10201488.html   (go语言格式化输出)
*/

func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v) // %T使用Go语法输出的值的类型

	v1 := "test"
	fmt.Printf("v is of type %T\n", v1)

	v2 := false
	fmt.Printf("v is of type %T\n", v2)

	v3 := 3.1415926
	fmt.Printf("v is of type %T\n", v3)

	v4 := 0.666 + 0.7i
	fmt.Printf("v is of type %T\n", v4)
}

/*
v is of type int
v is of type string
v is of type bool
v is of type float64
v is of type complex128
*/
