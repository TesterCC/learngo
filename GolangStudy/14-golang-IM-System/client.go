package main

import (
	"flag"
	"fmt"
	"net"
)

/*
v0.9 - 客户端实现 - 连接服务 p46
v0.9 - 客户端实现 - 命令行解析 p47
v0.9 - 客户端实现 - 菜单显示 p48

目标：
以终端客户端形式实现，实现nc链接server时的功能

client编译方法：
go build -o client client.go

server编译方法：
go build -o server main.go server.go user.go
*/

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
	}

	// 链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return nil
	}
	client.conn = conn

	// 返回对象
	return client

}

var serverIp string
var serverPort int

// useage: ./client -ip 127.0.0.1 -port 8888

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址（默认是127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口（默认是8888）")
}

func main() {

	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>> 服务器链接失败 ...")
	}
	fmt.Printf(">>>> 服务器链接成功，已连上：%s:%d \n", serverIp, serverPort)

	// 启动客户端的业务
	select {

	}
}
