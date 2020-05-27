package main

import "fmt"

func main() {
	numbers := [6]int{1, 3, 2, 5, 8, 4}

	// 简单遍历
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// 数组遍历
	maxi := -1
	maxValue := -1
	for i, v := range numbers {
		if v > maxValue {
			maxi, maxValue = i, v
		}
	}
	fmt.Println(maxi, maxValue)

	// 求和，可通过_省略变量
	// 不仅range，任何地方都可通过_省略边量
	// 如果只要i，可写成 for i:= range numbers
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	fmt.Println(sum)
}
