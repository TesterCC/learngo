package main

import (
	"fmt"
	"time"
)

func processData(val int) int {
	time.Sleep(1 * time.Second)
	return 2 * val

}

// 优化Go并发性能
// https://www.bilibili.com/video/BV1L52UYDEjT
func main() {
	input := []int{1, 2, 3, 4, 5}
	output := []int{}

	// record start time
	startTime := time.Now()

	for _, v := range input {
		result := processData(v)
		output = append(output,result)
	}

	// count cost time  // 5.02s
	fmt.Println(time.Since(startTime))
	fmt.Println(output)
}
