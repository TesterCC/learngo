package main

import (
	"fmt"
	"math"
)

/*

https://tour.go-zh.org/flowcontrol/6
https://go.dev/tour/flowcontrol/6


if 和简短语句

和 for 一样，if 语句可以在条件表达式前执行一个简短语句。

该语句声明的变量作用域仅在 if 之内。

（在最后的 return 语句处使用 v 看看。）

*/

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10), // 9
		pow(3, 3, 20), // 20
		pow(4, 4, 30), // 30
	)
}
