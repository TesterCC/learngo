package main

import "fmt"

/*
闭包是携带状态的函数，将函数内外连接起来。
通过闭包，可以读取函数内部的变量。
可以使用闭包封装私有状态，让它们常驻于内存当中。
闭包能够引用其作用域上部的变量并进行修改，被捕获到闭包中的变量将随着闭包的生命周期一直存在，闭包中的变量使闭包本身具备了存储信息的能力。
*/

// P60  3-6 用闭包的特性实现一个简单的计数器

// createCounter函数返回一了一个闭包，该闭包中封装了计数值initial，从外部代码根本无法直接访问该变量。
// 不同的闭包之间变量不会互相干扰，c1和c2两个计数器都是独立进行计数的。
func createCounter(initial int) func() int {
	if initial < 0 {
		initial = 0
	}

	// 引用initial创建一个闭包
	return func() int {
		fmt.Printf("initial in closure: %v\n", initial)
		initial++
		fmt.Printf("after initial++ in closure: %v\n", initial)
		// 返回当前计数
		return initial
	}
}

func main() {
	// 计数器1
	c1 := createCounter(1)
	fmt.Println(c1())
	fmt.Println(c1())

	// 计数器2
	c2 := createCounter(10)
	fmt.Println(c2())
	fmt.Println(c1())
}
