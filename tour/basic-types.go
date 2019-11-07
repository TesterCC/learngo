package main

import (
	"fmt"
	"math/cmplx"
)

/*

https://tour.golang.org/basics/11

Go的基本类型

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名  // 表示一个 Unicode 码点

float32 float64

complex64 complex128

*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// 占位符：%T相应值的类型的Go语法表示, %v相应值的默认格式
	fmt.Printf("Type: %T Vaule: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Vaule: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Vaule: %v\n", z, z)
}
