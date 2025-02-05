package main

import (
	"fmt"
	"sync"
	"time"
)

/*
P94  除了使用channel进行goroutine之间的通信和同步操作，还可以使用sync包下的并发工具

1. Mutex 互斥锁
sync.Mutex 互斥锁能够保证在同一个时间段内仅有一个goroutine持有锁，有且仅有一个goroutine访问资源
其它申请锁的goroutine将会被阻塞直到锁被释放，然后，重新争抢锁的持有权
*/

func main() {
	var lock sync.Mutex
	go func() {
		// add lock
		lock.Lock()
		defer  lock.Unlock()
		fmt.Println("func1 get lock at " + time.Now().String())
		time.Sleep(time.Second)
		fmt.Println("func1 release lock at " + time.Now().String())

	}()

	time.Sleep(time.Second /10)

	go func() {
		// add lock
		lock.Lock()
		defer  lock.Unlock()
		fmt.Println("func2 get lock at " + time.Now().String())
		time.Sleep(time.Second)
		fmt.Println("func2 release lock at " + time.Now().String())

	}()

	// 等待所有goroutine执行完毕
	time.Sleep(time.Second * 4)

	// 结果按严格顺序执行
}