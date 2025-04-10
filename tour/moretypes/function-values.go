package main

import (
	"fmt"
	"math"
)

/*
https://tour.go-zh.org/moretypes/24
https://go.dev/tour/moretypes/24


函数值

函数也是值。它们可以像其它值一样传递。
函数值可以用作函数的参数或返回值。

*/

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))      // 13
	fmt.Println(compute(hypot))    // 5
	fmt.Println(compute(math.Pow)) // 81
}
