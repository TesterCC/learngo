package main

import "time"

// https://www.bilibili.com/video/BV1io4y1b75M

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
}

// addr port 必填
func NewServer(addr string, port int) *Server {
	return &Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  time.Second,
		MaxConns: 10,
	}
}

// call
func main() {
	NewServer("127.0.0.1", 8080)
}
