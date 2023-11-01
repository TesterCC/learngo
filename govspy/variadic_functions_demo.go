package main

import "fmt"

// Variadic Functions  可变参数函数  somefunction(*args)
// http://govspy.peterbe.com/#variadic_functions
// ... 表示参数是可变的，可以传入零个或多个参数。这些参数会被作为一个切片（slice）传递给函数体内部的代码。在函数体中，你可以像操作普通切片一样操作这个参数。
func average(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func main() {
	//fmt.Println(average(1, 2, 3, 4))    // 2.5
	fmt.Println(average(1, 2, 3, 4, 5)) // 3
}
