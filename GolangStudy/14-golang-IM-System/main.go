package main

// 程序入口
/*
cd ~/14-golang-IM-System
go build -o server main.go server.go
./server

在Terminal用nc尝试连接，看服务器返回信息
nc 127.0.0.1 8888
 */

func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()   // main.go和server.go都属于main包，所以不用import了。
}
