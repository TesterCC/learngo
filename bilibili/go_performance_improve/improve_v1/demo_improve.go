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
// https://www.bilibili.com/video/BV1L52UYDEjT   07：54

func main() {
	input := []int{1, 2, 3, 4, 5}
	output := []int{}   // 共享资源，需要用锁保护

	// wait all go routine finish
	var wg sync.WaitGroup
	var mu sync.Mutex

	// record start time
	startTime := time.Now()

	for _, v := range input {

		wg.Add(1) // 计数器+1
		// 协程接收一个值v
		go func(val int) {
			// 等函数运行完后才运行这行代码
			defer wg.Done() // 计数器-1

			// 同一时间，只有一个协程能拥有这把锁，抢不到的继续等待
			//开始前加锁


			result := processData(v)  // 是不需要共享资源的，可以移动到锁的逻辑外

			mu.Lock()
			output = append(output, result)  // 这个才是真正的共享资源， 1.00s
			mu.Unlock()

		}(v)

	}

	// 不使用wg.Wait()会直接退出，不会等待协程运行；会阻塞等待直到计数器为0
	wg.Wait()

	// count cost time  // 5.02s
	fmt.Println(time.Since(startTime))
	fmt.Println(output)
}
