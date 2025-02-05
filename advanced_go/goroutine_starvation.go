package main

import (
	"fmt"
	"runtime"
)

// P420 Appendix 8.独占CPU导致其它Goroutine饿死
// Goroutine是协作式抢占调度，Goroutine本身不会主动放弃CPU
// 解决方法1：是在for循环加入runtime.Gosched()调度函数
// 解决方法2：通过阻塞的方式避免CPU占用

// 解决方法1：是在for循环加入runtime.Gosched()调度函数
func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i:=0; i< 10; i++ {
			fmt.Println(i)
		}
	}()

	for {
		runtime.Gosched()
	}
}
