package main

import (
	"fmt"
	"strconv"
)

// pdf P175-176 短变量声明中至少一个变量名是新的。新变量名被视为声明，现有名字被视为赋值。
func main() {
	a := 1
	b, a := 2, 3
	a, c := 4, 5
	fmt.Println(a, b, c) // 4 2 5

	// go允许对所有事物使用短变量声明语法
	e, err := strconv.ParseFloat("1.23", 64)
	f, err := strconv.ParseFloat("4.56", 64)
	fmt.Println(e, f, err)
}
