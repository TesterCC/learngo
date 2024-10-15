package main

import (
	"fmt"
	"time"
)

func processData(val int) int {
	time.Sleep(1 * time.Second)
	return 2 * val
}

// channel应该是最优版本，性能差距不大，单代码相对简洁

func main() {
	input := []int{1, 2, 3, 4, 5}
	output := make([]int, len(input))

	// 创建一个有缓冲的通道，缓冲大小为 len(input)
	ch := make(chan int, len(input))

	// record start time
	startTime := time.Now()

	for _, v := range input {
		go func(val int) {
			result := processData(val)
			ch <- result
		}(v)
	}

	// 从通道中接收结果并填充输出切片
	for i := range output {
		output[i] = <-ch
	}

	// count cost time  // 1.000s
	fmt.Println(time.Since(startTime))
	fmt.Println(output)
}
