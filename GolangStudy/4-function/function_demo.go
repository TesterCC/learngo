package main

import "fmt"

// 一个返回值
func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	return c
}

// 多个返回值, 形参匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 666, 999
}

// 多个返回值, 形参有名字
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("--- foo3 ---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// 赋值给有名称的变量，作为返回值
	r1, r2 = 1000, 2000
	return r1, r2
}

// 和foo3细微的返回参数设置不同
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("--- foo3 ---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// 赋值给有名称的变量，作为返回值
	r1, r2 = 1000, 2000
	return r1, r2
}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("haha", 999)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)

	ret1, ret2 = foo3("foo3", 333)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)

	ret1, ret2 = foo4("foo3", 444)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)
}
