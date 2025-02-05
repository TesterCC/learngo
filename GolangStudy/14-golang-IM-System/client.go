package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

/*
v0.9 - 客户端实现 - 连接服务 p46
v0.9 - 客户端实现 - 命令行解析 p47
v0.9 - 客户端实现 - 菜单显示 p48
v0.9 - 客户端实现 - 更新用户名 p49
v0.9 - 客户端实现 - 公聊模式 p50
v0.9 - 客户端实现 - 私聊模式 p51

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
	flag       int // 当前client的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
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

// 处理server回应的消息，直接显示到标准输出即可 goroutine
func (client *Client) DealResponse() {
	// 一旦client.conn有数据，就直接copy到stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, client.conn)

	//// 等价于  这单段代码有一定问题
	//for {
	//	buf := make([]byte, 4096)
	//	client.conn.Read(buf)
	//	fmt.Println(buf)
	//}

}

// 新增menu()方法，获取用户输入的模式
func (client *Client) menu() bool {
	var flag int // 用于接收用户输入

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>> 请输入合法范围内的数字")
		return false
	}

}

// 查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err: ", err)
		return
	}

}

// 私聊模式
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.SelectUsers()
	fmt.Println(">>>> 请输入聊天对象[用户名], exit退出：")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>> 请输入私聊消息内容，exit退出：")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			// 消息不为空则发送
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write err: ", err)
					break
				}
			}

			chatMsg = ""
			fmt.Println(">>>> 请输入私聊聊天内容，exit退出。")
			fmt.Scanln(&chatMsg)
		}

		client.SelectUsers()
		fmt.Println(">>>> 请输入聊天对象[用户名], exit退出：")
		fmt.Scanln(&remoteName)
	}

}

// 公聊模式
func (client *Client) PublicChat() {
	// 提示用户输入消息
	var chatMsg string

	fmt.Println(">>>> 请输入聊天内容，exit退出。")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 发给服务器
		// 消息不为空则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn Write err: ", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>> 请输入聊天内容，exit退出。")
		fmt.Scanln(&chatMsg)
	}

}

func (client *Client) UpdateName() bool {
	fmt.Println(">>>> 请输入用户名：")
	fmt.Scanln(&client.Name)

	if client.Name == "exit" {
		fmt.Println("请不要用系统关键词命名！")
		return false
	}
	sendMsg := "rename|" + client.Name + "\n"

	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err: ", err)
		return false
	}

	return true
}

func (client *Client) Run() {
	for client.flag != 0 {
		// 不断提示用户输入菜单，知道返回true为止
		for client.menu() != true {

		}

		// 根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			// 公聊模式
			//fmt.Println("公聊模式选择...")
			client.PublicChat() // fixme: P50 公聊有有空格输出时自动换行了
		case 2:
			// 私聊模式
			//fmt.Println("私聊模式选择...")
			client.PrivateChat()
		case 3:
			// 更新用户名
			//fmt.Println("更新用户名选择...")
			client.UpdateName()
		}
	}
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

	// 单独开启一个goroutine去处理server的回执消息
	go client.DealResponse()

	fmt.Printf(">>>> 服务器链接成功，已连上：%s:%d \n", serverIp, serverPort)

	// 启动客户端的业务
	//select {}
	client.Run()
}
