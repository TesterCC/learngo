package main

import "fmt"

// Variadic Functions  可变参数函数  somefunction(*args)
// http://govspy.peterbe.com/#variadic_functions

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
