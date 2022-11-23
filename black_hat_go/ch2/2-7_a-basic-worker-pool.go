package main

import (
	"fmt"
	"sync"
)

/*
29 Port Scanning Using a Worker Pool
2-7 A basic worker pool

为避免不一致，使用goroutine池来管理并发的执行。
使用for循环创建一定数量goroutines作为资源池。
然后，在main()所在的“线程”中使用channel来提供任务。
*/

// A worker function for processing work
// channel will be used to receiver work
// WaitGroup will be used to track when a single work item has been completed.
// 需要两个参数：int类型的channel和WaitGroup指针。channel用来接收工作，WaitGroup用于标记工作完成。
func worker(ports chan int, wg *sync.WaitGroup){
	for p := range ports{
		fmt.Println(p)    // print port
		wg.Done()
	}
}

// manage the workload and provide work to your worker(int, *sync.WaitGroup) function.
func main() {
	// 创建100个工人，int类型channel的新程序，并输出到屏幕上。使用WaitGroup阻塞执行。
	ports := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 0; i< cap(ports); i++ {
		go worker(ports, &wg)
	}

	for i:=1; i<=1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}
