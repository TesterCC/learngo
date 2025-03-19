package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
使用 Go 语言实现的指数退避策略（Exponential Backoff）代码，包含带抖动的退避和最大重试次数的限制

Python版本见：
Python3Scripts/dev_demo/exponential_backoff/retry_demo.py
*/

const (
	maxRetries = 5 // 最大重试次数
	baseDelay  = 1 // 初始延迟时间（秒）
)

// 模拟任务函数 launchTask 模拟执行任务，固定返回 false 表示任务失败
func launchTask() bool {
	fmt.Println("mock run task")
	return false
}

func main() {
	//rand.Seed(time.Now().UnixNano()) // 初始化随机数种子 go 1.20-
	rand.Float64() // 初始化随机数种子 go 1.20+
	for attempt := 0; attempt < maxRetries; attempt++ {
		result := launchTask()
		if result {
			break
		}
		if attempt == maxRetries-1 {
			panic("Max retries exceeded")
		}
		// 计算延迟时间（指数退避 + 抖动）
		delay := baseDelay * (1 << attempt)
		jitter := rand.Float64()
		totalDelay := float64(delay) + jitter
		//fmt.Printf("[D] retry delay time: %d\n", totalDelay)
		//fmt.Printf("[D] Retry delay time: %.2f seconds\n", float64(totalDelay))
		fmt.Printf("[D] Retry delay time: %v seconds\n", totalDelay)
		time.Sleep(time.Duration(totalDelay) * time.Second)
	}
}
