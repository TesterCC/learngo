package main

import "fmt"

/*
https://tour.go-zh.org/moretypes/7
https://go.dev/tour/moretypes/7

切片

每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。
在实践中，切片比数组更常用。

类型 []T 表示一个元素类型为 T 的切片。

切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：

a[low : high]

它会选择一个半闭半开区间，包括第一个元素，但排除最后一个元素。

左闭右开区间，包含左侧端点，但不包含右侧端点。

以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素，因为最后一个4要排除：

a[1:4]

*/

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // 包含切片primes中下标从 1 到 3 的元素，下标是从0开始的， 这里[]int可省略
	fmt.Println(s)
}
