package main

import (
	"fmt"
	"net"
	"time"
)

// same to 2-2_for-loop-port-scanner.go
// P26  然而大部分网站是拒绝响应的，只有几个常用测试站点可用 -- 单线程扫描，低效
//var Target = "scanme.nmap.org"
var Target = "testphp.vulnweb.com"

func main() {
	fmt.Println("Start to scanning...")
	for i:= 1; i <= 1024; i++ {
		address := fmt.Sprintf(Target+":%d", i)
		fmt.Println(address)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered.
			continue
		}
		conn.Close()
		fmt.Printf("[+] Port %d open\n", i)
		time.Sleep(1 * time.Second)
	}
}
