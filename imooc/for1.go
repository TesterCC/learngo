package main

import "fmt"

// 2-5 循环

/*
for的条件里不需要括号
for的条件里可以省略初始条件，结束条件，递增表达式
*/

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

}
