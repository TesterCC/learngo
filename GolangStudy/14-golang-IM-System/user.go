package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

// 创建一个用户API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String() // 获取到当前客户端链接的地址

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// https://www.bilibili.com/video/BV1gf4y1r79E?p=41  将server.go中的用户业务封装到user.go中
// 用户的上线业务
func (this *User) Online()  {
	// 用户上线，将用户加入到onlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.server.BroadCast(this, "已上线")
}

// 用户的下线业务
func (this *User) Offline() {
	// 用户下线，将用户从OnlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播当前用户下线消息
	this.server.BroadCast(this, "下线")

}

// 给当前User对应的客户端发送消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前在线用户有哪些
		this.server.mapLock.Lock()
		for _,user := range this.server.OnlineMap{
			onlineMsg := "[*]["  + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg)>7 && msg[:7] == "rename|" {
		// v0.6 修改用户名
		// 消息格式：rename|张三
		newName := strings.Split(msg, "|")[1]  // 作者写法
		//newName := strings.Split(msg, "|")[8:]  // 可能名字中也有|，故可以取[8:]

		// 判断name是否存在
		_,ok := this.server.OnlineMap[newName]
		if ok {
			// 说明查询成功，当前key存在
			this.SendMsg("[X] 当前用户名已被使用\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("您已更新用户名：" + this.Name  + "\n")
		}

	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式：to|张三|消息内容

		// 1 获取对方的用户名 // P45
	} else {
		// 正常的消息广播
		this.server.BroadCast(this, msg)
	}
}


// 监听当前User channel的方法，一旦有消息，就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		// 将当前读取的数据发送
		this.conn.Write([]byte(msg + "\n"))
	}

}
