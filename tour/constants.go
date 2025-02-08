package main

import "fmt"

/*

https://tour.go-zh.org/basics/15
https://go.dev/tour/basics/15

常量

常量的声明与变量类似，只不过使用 const 关键字。

常量可以是字符、字符串、布尔值或数值。e.g. character, string, boolean, or numeric values.

常量不能用 := 语法声明。

*/

const Pi = 3.14

func main() {

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

}
