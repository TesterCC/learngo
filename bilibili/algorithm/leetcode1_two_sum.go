package main

import "fmt"

// ref: https://leetcode.cn/problems/two-sum/description/

func twoSum(nums []int, target int) []int {
	// time complexity: O(N)
	// space complexity: O(N)
	hashTable := map[int]int{}

	for i, x := range nums {
		if j, ok := hashTable[target-x]; ok {
			return []int{j, i}
		}
		hashTable[x] = i

		fmt.Println("current hashTable: ", hashTable) // debug
	}
	return nil
}

func main() {
	//nums := []int{2, 7, 11, 15}
	//target := 9
	nums := []int{2, 3, 11, 15}
	target := 18

	ret := twoSum(nums, target)
	fmt.Println(ret)
}
