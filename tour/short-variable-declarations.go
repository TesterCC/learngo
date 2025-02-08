package main

import "fmt"

/*
https://tour.go-zh.org/basics/10
https://go.dev/tour/basics/10

短变量声明
在函数中，短赋值语句 := 可在隐式确定类型的 var 声明中使用。
函数外的每个语句都 必须 以关键字开始（var、func 等），因此 := 结构不能在函数外使用。
*/

func main() {
	var i, j int = 1, 2
	k := 3 // Attention: Outside a function, := construct is not available.
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
