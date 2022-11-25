package main

import (
	"fmt"
	"net"
	"sort"
)

/*
P30-31 Multichannel Communication
2-8 Port scanning with multiple channels

为避免不一致，使用goroutine池来管理并发的执行。
使用for循环创建一定数量goroutines作为资源池。
然后，在main()所在的“线程”中使用channel来提供任务。

2-7 issue：该扫描器不会按顺序检查端口，所以打印的端口是无序的。

2-8 modification and solution:
1.需要使用一个单独的线程将端口扫描的结果传回你的主线程，以便在打印前对端口进行排序。
2.可以完全删除对 WaitGroup 的依赖性，因为你将拥有另一种跟踪完成的方法。

例如：
如果扫描1024个端口就要向worker的管道中发送1024次，并且也将结果发送到主线程(main thread)1024次。
因为发送的任务数量要和收到的结果数量相同，因此程序就能知道何时关闭管道并随后关闭worker。

P.S.: 确实很快
*/

// worker(ports, results chan int)函数修改为接收两个管道；其余的逻辑一样，当前端口关闭就发送0值，当前端口开启就发送该端口值。
func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// 当前端口关闭就发送0值
			results <- 0
			continue
		}
		conn.Close()
		// 当前端口开启就发送该端口值
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
