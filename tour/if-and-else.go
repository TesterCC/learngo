package main

import (
	"fmt"
	"math"
)

/*

https://tour.go-zh.org/flowcontrol/7
https://go.dev/tour/flowcontrol/7

在 if 的简短语句中声明的变量同样可以在对应的任何 else 块中使用。

（在 main 的 fmt.Println 调用开始前，两次对 pow 的调用均已执行并返回其各自的结果。）

*/

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// here also can use var v
		fmt.Printf("v: %g >= lim: %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

// 在 main 的 fmt.Println 调用开始前，两次对 pow 的调用均已执行并返回其各自的结果。
func main() {
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
		pow2(4, 4, 30),
	)
}
