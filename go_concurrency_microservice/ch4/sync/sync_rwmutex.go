package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*
P95 sync.RWMutex即读写锁，将读锁和写锁操作分离开来，为了保证在同一时间段能够有多个goroutine访问同一资源。
读写锁需要满足的条件：
- 在同一时间段只能有一个goroutine获取到写锁
- 在同一时间段可以有任意多个goroutine获取到读锁
- 在同一时间段只能存在写锁或读锁（读锁和写锁互斥）
*/

// sync.RWMutex允许多读和单写
// 在写锁没有被获取时，所有read goroutine可以同时申请到读锁

var rwLock = sync.RWMutex{}

func main() {
	// get read lock
	for i := 0; i < 5; i++ {
		go func(i int) {
			rwLock.RLocker()
			defer rwLock.RLocker()
			fmt.Println("read func " + strconv.Itoa(i) + "get rlock at " + time.Now().String())
			time.Sleep(time.Second)
		}(i)
	}

	time.Sleep(time.Second / 10)

	// get write lock
	for i := 0; i < 5; i++ {
		go func(i int) {
			rwLock.Lock()
			defer rwLock.Unlock()
			fmt.Println("write func " + strconv.Itoa(i) + "get wlock at " + time.Now().String())
			time.Sleep(time.Second)
		}(i)

	}

	// 保证所有的goroutine执行结束
	time.Sleep(time.Second * 10)

}
