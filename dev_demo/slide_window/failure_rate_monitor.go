package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
实现使用滑动窗口记录一个窗口内的失败请求比率，超过 20% 就不再重试的代码。

Go 代码

    simulateRequest 函数模拟请求，随机返回成功或失败。
    slidingWindowFailureRate 函数使用切片实现滑动窗口，每次请求后更新窗口，计算失败率，若超过阈值则停止重试。
    在 main 函数中设置随机数种子，调用 slidingWindowFailureRate 函数进行测试。

Python代码

/Python3Scripts/dev_demo/slide_window/failure_rate_monitor.py

*/

// simulateRequest 模拟请求，随机返回成功或失败
func simulateRequest() bool {
	return rand.Float64() < 0.8
}

// slidingWindowFailureRate 实现滑动窗口记录失败请求比率 todo learn
func slidingWindowFailureRate(maxRetries, windowSize int, failureThreshold float64) {
	window := make([]bool, 0, windowSize)
	for attempt := 0; attempt < maxRetries; attempt++ {
		success := simulateRequest()
		if len(window) == windowSize {
			window = window[1:]
		}
		window = append(window, !success)
		var failureCount int
		for _, failed := range window {
			if failed {
				failureCount++
			}
		}
		failureRate := float64(failureCount) / float64(len(window))
		if failureRate > failureThreshold {
			fmt.Printf("Failure rate %.2f%% exceeds threshold, stopping retries.\n", failureRate*100)
			break
		}
		if success {
			fmt.Println("Request succeeded.")
			return
		}
	}
	fmt.Println("Max retries reached without success.")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	maxRetries := 10
	windowSize := 5
	failureThreshold := 0.2
	slidingWindowFailureRate(maxRetries, windowSize, failureThreshold)
}
