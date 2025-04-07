package main

import (
	"fmt"
	"time"
)

// https://www.bilibili.com/video/BV1io4y1b75M

type Server struct {
	Addr   string
	Port   int
	Config *Config
}

// 将可选参数放入配置对象中
type Config struct {
	Protocol string
	Timeout  time.Duration
	MaxConns int
}

// addr port 必填
func NewServer(addr string, port int, config *Config) *Server {
	if config != nil && config.Protocol == "" {
		config.Protocol = "tcp"
	}

	if config != nil && config.Timeout == 0 {
		config.Timeout = time.Second
	}

	if config != nil && config.MaxConns == 0 {
		config.MaxConns = 10
	}

	fmt.Println(config.Protocol, config.Timeout, config.MaxConns) // debug

	return &Server{
		Addr:   addr,
		Port:   port,
		Config: config,
	}
}

// call
func main() {
	// 缺点：不优雅
	NewServer("127.0.0.1", 8888, &Config{"udp", 20, 20})
	NewServer("127.0.0.1", 8889, &Config{Protocol: "ftp", Timeout: 30})
}
