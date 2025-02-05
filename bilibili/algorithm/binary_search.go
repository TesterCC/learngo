package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Binary Search 二分查找
- 输入：已排好序的集合
- 如要查找的元素在集合中：返回该元素位置（索引）
- 否认：返回空

- 二分查找：从中间开始查找，每次都能排除一半的元素
- 时间复杂度：O(log n)

ref:
https://www.bilibili.com/video/BV1Hg411G7dc
*/

func main() {
	// go 1.13 Digit separators  https://golang.org/doc/go1.13
	list := make([]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		list = append(list, i+1)
	}
	//fmt.Printf("list length is %d \n", len(list))  // 长度200万

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		v := rand.Intn(1_000_000-1) + 1
		fmt.Printf("针对%d进行二分查找：\n", v)
		idx := binarySearch(list, v)
		fmt.Printf("%d的索引位置是：[%d]\n", v, idx)
		fmt.Println("------------------------------------------")
	}
}

func binarySearch(list []int, target int) int {
	low := 0
	high := len(list) - 1

	step := 0

	for {
		step = step + 1
		if low <= high {
			mid := (low + high) / 2
			guess := list[mid]
			if guess == target {
				fmt.Printf("共查找了 %d 次\n", step)
				return mid
			}
			if guess > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}

}
