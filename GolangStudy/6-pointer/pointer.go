package main

import "fmt"

/*
8小时转职Golang工程师
ref: https://www.bilibili.com/video/BV1gf4y1r79E?p=10

*/


// 值传递，不能交换成功
func swap0(a int, b int)  {
	var temp int
	temp = a
	a = b
	b = temp
}

// 指针传递/引用传递，可以交换成功
func swap(pa,pb *int)  {
	var temp int
	temp = *pa // temp = main::a
	*pa = *pb  // main::a = main::b
	*pb = temp // main::b = temp
	
}

// 经典例子，交换2个数据
func main() {
	var a int = 10
	var b int = 20

	// swap
	//swap0(a,b)   // a =  10 b =  20, 值传递，不能交换成功
	swap(&a, &b)   // a =  20 b =  10， 注意这里是传变量的地址

	fmt.Println("a = ", a, "b = ", b)

	// p是指针类型
	var p *int
	p = &a
	fmt.Println("&a: ", &a)  // 0xc00000a0c0
	fmt.Println("p: ", p)    // 0xc00000a0c0

	// 二级指针
	var pp **int   // 二级指针(go语言中比较少见)，用于存一级指针的地址
	pp = &p
	fmt.Println("&p: ", &p)  // 0xc000006030
	fmt.Println("pp: ", pp)  // 0xc000006030
}

