/*

https://tour.go-zh.org/methods/5
https://go.dev/tour/methods/5


指针与函数

现在我们要把 Abs 和 Scale 方法重写为函数。

同样，先试着移除掉第 16 的 *，你能看出程序行为改变的原因吗？ 要怎样做才能让该示例顺利通过编译？
*/

package main

import (
	"fmt"
	"math"
)

type Vertex4 struct {
	X, Y float64
}

// 计算其模长（欧几里得范数）
func Abs(v Vertex4) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 普通函数，但它使用的是 值传递，不会修改 v 的原值。
func Scale(v Vertex4, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex4{3, 4}
	Scale(v, 10)
	fmt.Println(Abs(v))
}
