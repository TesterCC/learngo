package main

import (
	"fmt"
	"time"
)

/*

P59

匿名函数： 一种没有函数名的函数，即定义即使用

匿名函数没有函数名，只有函数体，只有在被调用的时候才会被初始化。
匿名函数一般被当作一种 类型 赋值给函数类型的变量，经常被用作回调函数。

闭包：作为一种携带状态的函数。 可以理解为"对象"， 它同时具备状态和行为

*/

func main() {

	// 匿名函数
	func(name string) {
		fmt.Println("My name is", name)
		fmt.Printf("My name also is %v.\n", name)
	}("Lily")   // 在声明匿名函数后，在其后加上调用的参数列表，即可对匿名函数进行调用。

	// 还可以将 匿名函数 赋值给 函数类型的变量，用于多次调用或者求值
	currentTime := func(){
		fmt.Println(time.Now())
	}
	//调用匿名函数
	currentTime()
}
