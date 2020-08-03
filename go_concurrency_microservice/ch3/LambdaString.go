package main

import "fmt"

/*
匿名函数： 一种没有函数名的函数，即定义即使用

匿名函数没有函数名，只有函数体，只有在被调用的时候才会被初始化。
匿名函数一般被当作一种 类型 赋值给函数类型的变量，经常被用作回调函数。

P59-60  3-5 使用回调函数处理字符串
*/

// 定义函数：它接受 string 和 匿名函数的参数输入，然后使用匿名函数对string进行处理。
// 匿名函数作为 回调函数 用于对传递的字符串进行处理

func proc(input string, processor func(str string)) {
	//调用匿名函数
	processor(input)
}

func main() {
	proc("Tester", func(str string) {
		for _, v := range str {
			fmt.Printf("%c\n", v)
		}
	})

}
