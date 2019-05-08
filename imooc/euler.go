package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//2-2 内建变量类型

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

func main() {
	euler()
	triangle()
}
