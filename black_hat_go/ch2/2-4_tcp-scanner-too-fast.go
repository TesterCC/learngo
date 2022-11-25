package main

import (
	"fmt"
	"net"
)

// P26-27  use goroutine, 2-4 A scanner that works too fast

/*
test cmd: time go run 2-4_tcp-scanner-too-fast.go

time是Unix/Linux内置多命令，使用时一般不用传过多参数，直接跟上需要调试多程序即可。
- real：从程序开始到结束，实际度过的时间；
- user：程序在用户态度过的时间；
- sys：程序在内核态度过的时间。

一般情况下 real >= user + sys，因为系统还有其它进程(切换其他进程中间对于本进程会有空白期)。

*/

func main() {
	//var target = "testphp.vulnweb.com"
	//并发扫描的最本质的做法是将Dial(network, address string)封装在一个goroutine中去调用。
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			//address := fmt.Sprintf(target+":%d", i)
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			//address := fmt.Sprintf("testphp.vulnweb.com:%d", j)
			fmt.Println(address)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
			// Attention: 不能得到精准port信息，因为主协程没有等待连接完成。
			// 每一个连接分配了一个goroutine，main函数所在的goroutine并不知道要等待连接。因此，代码执行完for循环后就退出了，这要比代码里的网络通信快多了。也就无法得到正确的结果了。
		}(i)
	}
}
