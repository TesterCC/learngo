package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

/*
TCP sender
smoke testing with ~/debugger/tcp_receiver.py

用flag来传参，而不是用os.Args

gopl P25  flag example
ref: https://www.bilibili.com/video/BV1WF411u77z

go build tcp_sender_simple2.go
./tcp_sender_simple2 -ip 192.168.80.129 -port 19999

这个Go语言客户端程序做了以下几件事情：
使用net.Dial尝试与服务器中的TCP服务建立连接。这里使用了192.168.80.129:9999，表示192.168.80.129的9999端口。
发送一条消息给服务器。消息以换行符\n结束，这假设服务器将换行符视为消息结束的标志。
使用带缓冲的reader来读取直到第一个换行符的字符串，这样可以确保读取整个响应，而不仅仅是一个固定大小的块。
打印服务器的响应，并在交互结束后关闭TCP连接。
请确保在运行此客户端程序之前，服务器程序已经启动并监听在预期的端口上。如果服务器处于其他地址或端口，请相应地调整server变量的值。
*/


var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "目标服务器IP地址（默认是127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "目标服务器端口")
}

func main() {
	// 定义服务器地址
	//server := "192.168.80.129:19999"

	// 命令行解析
	flag.Parse()

	// 拼接变量
	server := fmt.Sprintf("%s:%d", serverIp, serverPort)
	fmt.Println("[D] Send data to: ", server)

	// 连接到服务器
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("连接到服务器", server)

	// 发送数据
	message := "Hello from client!"
	_, err = conn.Write([]byte(message + "\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("客户端发送：", message)

	// 读取响应
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("读取服务器响应出错：", err)
		return
	}

	fmt.Printf("服务器响应： %s", response)

	// 完成后关闭连接
	fmt.Println("客户端程序结束")
}
