package main

import (
	"fmt"
	"time"
)

// https://www.bilibili.com/video/BV1io4y1b75M
// 基于闭包函数实现功能选项模式
// 这也是go语言中主流的封装模式 recommend 1

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
}

// 对应的闭包函数入参为Server指针
type Option func(server *Server)

// 针对server的可选参数实现WithXX方法
// WithXX方法内部实现参数赋值
func WithProtocol(protocol string) Option {
	return func(s *Server) {
		s.Protocol = protocol
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func WithMaxConns(maxConns int) Option {
	return func(s *Server) {
		s.MaxConns = maxConns
	}
}

func NewServer(addr string, port int, options ...Option) *Server {
	// 对server初始化默认值后
	serv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  time.Second,
		MaxConns: 10,
	}

	// 遍历可变参数，执行对应的闭包函数进行赋值
	for _, option := range options {
		option(&serv)
	}

	fmt.Println(serv) // debug

	return &serv

}

func main() {
	// 功能选项方案-基于闭包函数
	NewServer("127.0.0.1", 9999)
	NewServer("127.0.0.1", 9000, WithProtocol("UDP"))
	NewServer("127.0.0.1", 8888, WithProtocol("FTP"), WithMaxConns(30))
}
