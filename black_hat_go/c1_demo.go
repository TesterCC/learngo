package main

import "fmt"

/*
P6-7  go run执行后不会生成独立的二进制文件，要用

生成可执行的二进制文件：
go build xxx.go
./xxx

默认go build包含debug信息和符号表，减少生成二进制文件大小约30%的方法
go build -ldflags "-w -s" xxx.go
*/

func main() {
	fmt.Println("Hello, Black Hat Gophers!")
}
