package main

import (
	"fmt"
	"math/big"
)

/*

https://tour.go-zh.org/basics/16
https://go.dev/tour/basics/16

数值常量

数值常量是 高精度的 值。

一个未指定类型的常量 由上下文来决定 其类型。

再试着一下输出 needInt(Big) 吧。

（int 类型可以存储最大 64 位的整数，根据平台不同有时会更小。）
*/

const (
	// 将 1 左移 100 位来创建一个非常大的数字 (1向左移动100位), 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	bigInt := new(big.Int)
	bigInt.Lsh(big.NewInt(1), 100) // 1 << 100

	//fmt.Printf("origin big is %d\n", Big)
	// 在 Go 语言中，int 类型的位数是有限的（通常是 32 位或 64 位），而 1 << 100 是一个非常大的数（1267650600228229401496703205376），远远超过了 int 类型的表示范围，因此会导致溢出错误。
	fmt.Printf("origin big is %v\n", bigInt) // 使用 "math/big" 可以处理任意大小的整数
	fmt.Printf("origin small is %d\n", Small)

	fmt.Println(needInt(Small))   // 21
	fmt.Println(needFloat(Small)) // 0.2
	fmt.Println(needFloat(Big))   // 1.2676506002282295e+29
}
