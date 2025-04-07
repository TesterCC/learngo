package main

import (
	"fmt"
	"time"
)

// https://www.bilibili.com/video/BV1io4y1b75M
// 基于接口实现功能选项模式
// 比基于闭包的方案更复杂，但优点是可扩展性和灵活性更强

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
}

// 定义一个包含未导出方法的Option接口
type Option interface {
	// 用于赋值参数
	apply(*Server)
}

// 对应的可选参数定义一个别名，这里是 protocolOption
type protocolOption string

// 别名类型实现Option接口
func (c protocolOption) apply(s *Server) {
	s.Protocol = string(c)
}

// 该方法供外部调用
func WithProtocol(protocol string) Option {
	return protocolOption(protocol)
}

// 其它的类似

func NewServer(addr string, port int, options ...Option) *Server {
	// 对server初始化默认值后
	serv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  time.Second,
		MaxConns: 10,
	}

	// 遍历可变参数，调用apply()进行赋值
	for _, option := range options {
		option.apply(&serv)
	}

	fmt.Println(serv) // debug

	return &serv

}

func main() {
	// 基于接口实现功能选项模式
	NewServer("127.0.0.1", 9999)
	NewServer("127.0.0.1", 9000, WithProtocol("UDP"))
}
