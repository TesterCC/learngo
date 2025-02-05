package main

import "fmt"

// https://leetcode-cn.com/problems/running-sum-of-1d-array/

var arr2 = []int{1, 2, 3, 4}   // 注意go的类型约束

// 这个解法思路容易理解一些
func runningSum2(nums []int) []int {
	sum := 0
	for i,_:=range nums {
		sum += nums[i]
		nums[i] = sum

		fmt.Println(nums)
	}
	return nums
}

func main() {
	runningSum2(arr2)
}
