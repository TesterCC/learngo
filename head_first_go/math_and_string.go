package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	math.Floor(2.75)               // 浮点数向下取整
	strings.Title("head first go") // 转换单词为首字母大写

	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first go"))

	fmt.Println('A', '\n')   // rune 单引号， A 数字代码 65 ; \n 数字代码 10
	fmt.Println("A")   // 字符 A
}
