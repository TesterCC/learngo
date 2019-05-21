package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//2-2 内建变量类型
/*
变量定义要点：
1.变量类型写在变量名之后
2.编译器可推测变量类型
3.没有char，只有rune (32 bits)
4.原生支持复数类型
*/



//验证欧拉公式
func euler() {
	c := 3 + 4i //define complex
	fmt.Println(cmplx.Abs(c))

	//println(cmplx.Pow(math.E, 1i* math.Pi)+1)   // method 1
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1) // mehtod 2 better

	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

//验证强制类型转换
func triangle() {
	var a, b int = 3,4
	//var c int = math.Sqrt(a*a + b*b)  //会提示错误
	var c int = int(math.Sqrt(float64(a*a + b*b)))    //浮点数容易失去精度  PS:基本思路：将浮点数转换成整数进行计算，然后再将整数结果转回正确的浮点数。
	fmt.Println(c)
}

//2-3 常量与枚举
// const数值可作为各种类型使用
const filename_outer = "def.txt"
const (
	test_name = "Go"
	version2 = 1.1
)

func consts(){
	const filename = "abc.txt"
	const a, b = 3, 4
	const version float32 = 1.01
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c, filename_outer)
	fmt.Println(test_name, version, version2)
}

//使用常量定义枚举类型
func enums(){
	const(
		//cpp = 0
		//java = 1
		//python = 2
		//golang = 3
		//php = 4
		//javascript = 5
		cpp = iota  //表示这组是自增值
		_    // 占位，算在自增序号中
		python
		golang
		php
		javascript
	)

	// b, kb, mb, gb, tb, pb
	const(
		b = 1 << (10*iota)  // 左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang, php)
	fmt.Println(b,kb,mb,gb,tb,pb)
}



func main() {
	euler()
	triangle()
	consts()
	enums()
}
