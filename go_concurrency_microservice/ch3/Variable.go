package main

import "fmt"

// P42

func main() {
	var a int = 100

	var b = "100"

	c := 0.17

	fmt.Printf("a value is %v, type is %T\n", a, a)
	fmt.Printf("b value is %v, type is %T\n", b, b)
	fmt.Printf("c value is %v, type is %T\n", c, c) // 为了提高精度，默认推导为float64

	x, y := 1, 2
	x, y = y, x
	fmt.Printf("x value is %v, y value is  %v.\n", x, y)
}
