package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	return c
}

func main() {
	c := foo1("abc", 777)
	fmt.Println("c = ", c)
}
