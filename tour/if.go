package main

import (
	"fmt"
	"math"
)

/*
https://tour.go-zh.org/flowcontrol/5
https://go.dev/tour/flowcontrol/5

if 判断

Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( )，而大括号 { } 则是必须的。
*/

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4)) // 1.4142135623730951 2i
}
