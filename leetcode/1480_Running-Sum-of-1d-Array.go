package main

import "fmt"

// https://leetcode-cn.com/problems/running-sum-of-1d-array/

var arr = []int{1, 2, 3, 4}   // 注意go的类型约束

func runningSum(nums []int) []int {
	for i := range nums {
		if i != 0 {
			nums[i] = nums[i-1] + nums[i]
		}
		fmt.Println(nums)
	}
	return nums
}


func main() {
	//fmt.Println(arr)
	runningSum(arr)

}
