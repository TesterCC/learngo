package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// 服务端 Server基本的listen操作

type Server struct {
	Ip   string
	Port int

	// 在线用户的列表 online user map ,  key当前用户名，value当前用户对象
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
	user := NewUser(conn, this)

    // user online
	user.Online()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)
	

	// 接收客户端client发送的消息
	// important 单独goroutine处理当前套接字的读请求
	go func() {
		buf := make([]byte, 4096)
		for {
			n,err := conn.Read(buf)
			// 读出0表示客户端是合法关闭
			if n == 0 {
				//this.BroadCast(user, "下线")
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err: ", err)
				return
			}

			// 提取用户的消息(去除'\n')
			msg := string(buf[:n-1])

			// 用户针对msg进行消息处理
			user.DoMessage(msg)

			// 用户的任意消息，代表当前用户是一个活跃的用户
			isLive <- true
		}
	}()

	for {
		// 保持当前handler阻塞 goroutine不死
		select {
		case <- isLive:
			// 当前用户是活跃的，应该重置定时器
			// 不做任何事情，为了要激活select，更新下面的定时器
        // timeout 20s
		case  <- time.After(time.Second * 20):
			// 代码逻辑进入这里说明已经超时
			// 将当前的User强制关闭
			user.SendMsg("你超时被踢了")

			// 销毁用户资源
			close(user.C)

			// 关闭连接
			conn.Close()

			// 退出当前Handler
			return // runtime.Goexit()
		}
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
