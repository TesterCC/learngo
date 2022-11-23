package main

import (
	"fmt"
	"net"
)

// P25  2-3 扫描scanme.nmap.org的1024个端口
//var target = "172.16.12.10"

func main() {
	for i := 1; i < 10; i++ {
		//address := fmt.Sprintf("target"+":%d", i)  // 变量拼接
		address := fmt.Sprintf("scanme.nmap.org:%d", i) // 变量拼接
		//fmt.Println(address)

		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered.
			//fmt.Println("Connection successful")
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)

		//time.Sleep(2 * time.Second)   // 增加延时
		//time.Sleep(1500 * time.Millisecond) // 1.5 seconds
	}
}
