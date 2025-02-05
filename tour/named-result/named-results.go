package main

import "fmt"

/*
https://tour.go-zh.org/basics/7

带名字的返回值

Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。

返回值的命名应当能反应其含义，它可以作为文档使用。

没有参数的 return 语句会直接返回已命名的返回值，也就是「裸」返回值。

裸返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。

*/

func main() {
	fmt.Println(split(17))

}

func split(sum int) (x, y, z int) {
	x = sum * 4 / 9    // 7.5 = 7
	y = sum - x        // 17-7=10
	z = (sum - 7) * 10 // 100
	return             // don't use "naked" return in longer functions. 即使 return 语句后面没有显式地返回值，Go语言也会自动返回这些命名返回值的当前值。
}
