package main

import "fmt"

/*
4种变量的声明方式
*/

// 声明全局变量  方法一、二、三
// 声明局部变量  方法四

func main() {
	// 方法一：声明一个变量 默认的值时0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	// 方法二：声明一个变量，初始化一个值
	var b = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb = "abcd"
	fmt.Printf("type of bb = %T\n", bb)

	// 方法三：在初始化的时候，可以省去数据类型，通过值自动匹配当前变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "efgh"
	fmt.Printf("type of cc = %T\n", cc)

	// 方法四：(常用方法)省去var关键字，直接自动匹配 (在函数体内，)
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := "xyz"
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	// 声明多个变量
	var xx, yy = 100, "test"
	fmt.Println("xx = ", xx, ", yy = ", yy)
}
