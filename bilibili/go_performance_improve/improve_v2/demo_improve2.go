package main

import (
	"fmt"
	"sync"
	"time"
)

func processData(val int) int {
	time.Sleep(1 * time.Second)
	return 2 * val

}

// 优化Go并发性能
// https://www.bilibili.com/video/BV1L52UYDEjT
// 思路：避免使用迸发和共享资源

func main() {
	input := []int{1, 2, 3, 4, 5}
	output := make([]int, len(input)) // initial [0,0,0,0,0]

	// wait all go routine finish
	var wg sync.WaitGroup

	// record start time
	startTime := time.Now()

	for i, v := range input {

		wg.Add(1) // 计数器+1
		// 协程接收一个值v
		go func(val int, p *int) {
			// 等函数运行完后才运行这行代码
			defer wg.Done() // 计数器-1

			result := processData(val)

			// *p 取值
			*p = result

		}(v, &output[i])

	}

	// 不使用wg.Wait()会直接退出，不会等待协程运行；会阻塞等待直到计数器为0
	wg.Wait()

	// count cost time  // 1.000s
	fmt.Println(time.Since(startTime))
	fmt.Println(output)
}
