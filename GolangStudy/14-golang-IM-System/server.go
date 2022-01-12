package main

import (
	"fmt"
	"net"
	"sync"
)

// 服务端 Server基本的listen操作

type Server struct {
	Ip   string
	Port int

	// online user map ,  key当前用户名，value当前用户对象
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 创建一个Server的接口   返回Server指针
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送到全部的在线User
func (this *Server) ListenMessager() {
	for {
		msg := <- this.Message

		// 将msg发送给全部在线User
		this.mapLock.Lock()
		for _,cli := range this.OnlineMap{
			cli.C <- msg
		}
		this.mapLock.Unlock()

	}

}


// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string){
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	//当前连接的业务
	//fmt.Println("连接建立成功")

	// 创建用户
	user := NewUser(conn)

	// 用户上线，将用户加入到onlineMap中
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	//广播当前用户上线消息
	this.BroadCast(user, "已上线")

	// 保持当前handler阻塞
	select {

	}

}

// 启动服务器的接口
// 当前类的成员方法
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	// close listen socket
	defer listener.Close()

	// 启动监听Message的goroutine
	go this.ListenMessager()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err: ", err)
			continue
		}

		// do handler
		go this.Handler(conn) // goroutine 异步协程处理
	}
}
