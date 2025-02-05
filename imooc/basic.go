package main

import "fmt"

//2-1 变量定义

//函数外定义变量 不能用 :=, 要用var or func
//不是全局变量，是包内部变量
//var aa = 3
//var ss = "shadows"
//var bb = true

//可以使用var()集中定义变量
var (
	aa = 3
	ss = "shadows"
	bb = true
)

//函数内定义变量  定义变量使用var关键字可放在函数内、或直接放在包内
//定义变量 默认为0
func variableZeroValue() {
	var a int
	var s string
	//fmt.Println(a,s)   //无法对打印""
	fmt.Printf("%d %q\n", a, s)
}

//定义变量 设定初始默认值
func variableInitialValue() {
	var a, b int = 7, 8 //可以赋两个变量的初始值
	var s string = "abc"
	var s1, s2 string = "Learn", "Go"
	fmt.Println(a, b, s, s1, s2)
}

//定义变量 让编译器自动决定类型
func variableTypeDeduction() {
	//不规定类型的变量可以写在一行
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

//定义变量 语法糖  常用  := 仅能在函数内使用
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5 //修改变量
	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
}
