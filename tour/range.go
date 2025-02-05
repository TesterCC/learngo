package main

import "fmt"

/*

http://127.0.0.1:3999/moretypes/16
https://tour.go-zh.org/moretypes/16

The range form of the for loop iterates over a slice or map.

for 循环的 range 形式可遍历切片或映射。

当使用 for 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。

*/

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

func main() {

	// 第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}
