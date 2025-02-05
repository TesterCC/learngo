package main

import (
	"fmt"
	"net"
	"time"
)

/*
白帽子安全开放实战 P17 2.1.1

TCP全链接扫描原理：调用socket的connect函数连接到目标IP的特定端口上，如果连接成功，则说明端口是开放的

实现对多IP的扫描，建议引入一个第三方包，该包实现了对多种风格IP的解析。
go get -u github.com/malfunkt/iprange
*/

// 最简单理解版本，但只能检测一个IP的一个端口
// Dial 与 DialTimeout 的区别是后者增加了超时时间，无论想创建什么协议的连接，都只需要调用这2个函数即可
// 最简单的 TCP全连接 端口扫描器，一次仅能检测一个IP的一个端口。
func Connect(ip string, port int) (net.Conn, error) {
	//
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ip, port), 2*time.Second)

	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()
	return conn, err
}

func main() {
	var ip = "10.0.4.148"
	//var port = 80
	var port = 5000
	_, err := Connect(ip, port)
	if err != nil {
		fmt.Printf("ip: %v, port: %v is closed\n", ip, port)
		return
	}
	fmt.Printf("ip: %v, port: %v is open\n", ip, port)

}
