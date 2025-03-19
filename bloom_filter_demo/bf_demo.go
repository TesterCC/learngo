package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/bits-and-blooms/bloom/v3"
)

// ref: https://github.com/bits-and-blooms/bloom
// python version： ～/Python3Scripts/bloom_filter_demo/bf_demo.py
const (
	// 预计插入的元素数量 10万
	capacity = 100000
	// 期望的误判率
	errorRate = 0.001
	// 模拟的抢券请求次数，100万
	requests = 1000000
	// 模拟的用户数量 10万
	users = 100000
	// 抢券成功率
	successRate = 0.1
)

func main() {
	// 初始化布隆过滤器
	filter := bloom.NewWithEstimates(capacity, errorRate)
	var wg sync.WaitGroup
	// 创建一个并发安全的互斥锁，用于保护布隆过滤器的操作
	var mu sync.Mutex

	// 模拟 100 万次抢券请求
	for i := 0; i < requests; i++ {
		wg.Add(1)
		userID := fmt.Sprintf("user_%d", i%users)
		go grabCoupon(userID, filter, &wg, &mu)
	}

	// 等待所有协程完成
	wg.Wait()
}

// grabCoupon 模拟用户抢券操作
func grabCoupon(userID string, filter *bloom.BloomFilter, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	// 加锁，确保对布隆过滤器的操作是线程安全的
	mu.Lock()
	// 检查用户是否已经抢到券
	if filter.TestString(userID) {
		fmt.Printf("用户 %s 已经抢到券，本次请求被拒绝。\n", userID)
		mu.Unlock()
		return
	}
	mu.Unlock()

	// 模拟抢券逻辑，简单随机判断是否抢券成功
	rand.Seed(time.Now().UnixNano())
	if rand.Float64() < successRate {
		// 加锁，确保对布隆过滤器的操作是线程安全的
		mu.Lock()
		// 再次检查用户是否已经抢到券，避免并发问题
		if !filter.TestString(userID) {
			// 抢券成功，将用户 ID 插入布隆过滤器
			filter.AddString(userID)
			fmt.Printf("用户 %s 抢券成功！\n", userID)
		}
		mu.Unlock()
	} else {
		fmt.Printf("用户 %s 抢券失败。\n", userID)
	}
}
