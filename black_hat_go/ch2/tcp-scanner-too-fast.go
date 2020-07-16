package main

import (
	"fmt"
	"net"
)

// P26-27  use goroutine
func main() {
	var target = "testphp.vulnweb.com"

	for i:=1;i<=1024;i++ {
		go func(j int){
			address := fmt.Sprintf(target+":%d", i)
			fmt.Println(address)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n",j)  // 不能得到精准port信息，因为主协程没有等待连接完成
		}(i)
	}
}
