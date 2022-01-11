package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

// 创建一个用户API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String() // 获取到当前客户端链接的地址

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}
	return user
}


// 监听当前User channel的方法，一旦有消息，就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <- this.C

		// 将当前读取的数据发送
		this.conn.Write()  // https://www.bilibili.com/video/BV1gf4y1r79E?p=39   7:46
	}
	
}

