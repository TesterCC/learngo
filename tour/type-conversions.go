package main

import (
	"fmt"
	"math"
)

/*
https://tour.go-zh.org/basics/13
https://go.dev/tour/basics/13

类型转换

表达式 T(v) 将值 v 转换为类型 T。

一些数值类型的转换：

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
或者，更加简短的形式：

i := 42
f := float64(i)
u := uint(f)

与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换。
试着移除例子中的 float64 或 uint 的类型转换，看看会发生什么。

*/

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z, f)
	fmt.Printf("x type is: %T\n", x)
	fmt.Printf("y type is: %T\n", y)
	fmt.Printf("z type is: %T\n", z)
	fmt.Printf("f type is: %T\n", f)
}

/*
3 4 5 5
x type is: int
y type is: int
z type is: uint
f type is: float64
*/
