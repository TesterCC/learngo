package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// ref: https://www.jb51.net/article/212967.htm
/*
实现原理：发送socket连接(IP+端口)，如果能连接成功，说明目标主机开放了某端口。
当要大量扫描端口时，需要写并发编程。

注意：
使用net库的Dial函数做为socket客户端，需要注意的是要设置超时时间，因为若主机不存在，或目标端口是关闭的，往往需要花费数秒才返回错误，这样扫描大量端口时效率会极其低下。
在go中可以使用net.Dialer结构体设置超时时间，然后在调用Dial方法。

测试可用
*/

var wg sync.WaitGroup

// 定义了两个通道，分别用于生产端口和限制连接数，如果不限制连接数，容易被对方检测到或导致对方服务器不能正常运行。

// 生产端口
var port = make(chan int, 10)

// 限制并发数
var connect = make(chan string, 5)

func Run(address string, start, end int) {

	go func() {
		for i := start; i <= end; i++ {
			port <- i
		}
	}()

	go func() {
		// 消费端口
		for p := range port {
			// 往通道写入目标地址，超过限制并发数会阻塞
			connect <- fmt.Sprintf("%s:%d", address, p)
		}
	}()

	go Connect()

}

func Connect() {
	// 并发请求
	for target := range connect {
		// 设置超时时间
		// 使用net库的Dial函数做为socket客户端，需要注意的是要设置超时时间，因为若主机不存在，或目标端口是关闭的，往往需要花费数秒才返回错误，这样扫描大量端口时效率会极其低下。在go中可以使用net.Dialer结构体设置超时时间，然后在调用Dial方法
		// 当然，也可以改用 net.DialTimeout
		d := net.Dialer{Timeout: time.Second}
		dial, err := d.Dial("tcp", target)
		if err == nil {
			fmt.Printf("%s open, success.\n", target)
			dial.Close()
		} else {
			fmt.Printf("%s failed.\n", target)
		}
		wg.Done()
	}
}

func main() {
	var start, end int
	var address string
	fmt.Println("[+] Task Launch, please input target info...")
	fmt.Printf("请输入目标IP：> ")
	fmt.Scan(&address)
	fmt.Printf("请输入起始端口：> ")
	fmt.Scan(&start)
	fmt.Printf("请输入结束端口：> ")
	fmt.Scan(&end)
	// waitGroup
	wg.Add(end - start + 1)
	Run(address, start, end)
	wg.Wait()
	fmt.Println("[-] Task Finish!")
}

//  实测：探测结果不精准