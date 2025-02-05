package main

import "fmt"

/*
https://github.com/greyireland/algorithm-pattern/blob/master/introduction/golang.md
 */

func main() {

	// 栈 stack
	fmt.Println("1.Stack Demo")
	// 创建栈
	stack := make([]int, 0)
	// push压入
	stack = append(stack, 10)
	// pop弹出
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	// 检查栈空
	fmt.Println(len(stack) == 0)

	//print v
	fmt.Println(v)

}
