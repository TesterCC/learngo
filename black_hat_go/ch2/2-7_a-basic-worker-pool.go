package main

import (
	"fmt"
	"sync"
)

/*
P29 Port Scanning Using a Worker Pool
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
	// 在worker(int, *sync.WaitGroup)函数中使用range持续地从ports管道消费数据，直到关闭channel才退出循环。
	for p := range ports{
		fmt.Println(p)    // print port, maybe instead of task
		wg.Done()
	}
}

// manage the workload and provide work to your worker(int, *sync.WaitGroup) function.
func main() {
	// 创建100个工人，int类型channel的新程序，并输出到屏幕上。使用WaitGroup阻塞执行。
	/*
	首先，使用make()创建channel，第二个参数100是设置channel的缓存大小，缓存的意思是不需要等待消费掉channel里的值就能往channel里写入。
	带有缓存的channel是处理多个消费者和生产者的理想产物。此处，channel的容量设为100，即生成者发送100个值后才会阻塞。
	这能稍微提高点性能，因为所有任务可以立即执行。
	 */
	ports := make(chan int, 100)
	var wg sync.WaitGroup

	// 使用for循环启动所需的worker数量
	// 在main()函数中依次遍历端口，将端口通过ports管道发送给worker。
	for i := 0; i< cap(ports); i++ {
		go worker(ports, &wg)
	}

	for i:=1; i<=1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	// 所有任务完成后
	close(ports)

	// 输出的端口结果是无序的
}
