package main

import (
	"fmt"
	"net"
)

// P25-26  2-3 扫描scanme.nmap.org的1024个端口
//var target = "172.16.12.10"

func main() {
	for i := 1; i < 1024; i++ {
		//address := fmt.Sprintf("target"+":%d", i)  // 变量拼接
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		//fmt.Println(address)

		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered.
			//fmt.Println("Connection successful")
			continue
		}
		// 如果连接成功，还应该添加些逻辑来关闭连接；这样连接就不会一直打开。
		conn.Close()
		fmt.Printf("%d open\n", i)

		//time.Sleep(2 * time.Second)   // 增加延时
		//time.Sleep(1500 * time.Millisecond) // 1.5 seconds
	}
	fmt.Println("[-] Scanner working end...")
}
