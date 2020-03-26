package main

import "fmt"

/*
P6-7  go run执行后不会生成独立的二进制文件，要用下面的命令生成可执行的二进制文件：
go build xxx.go
./xxx

默认go build包含debug信息和符号表，减少生成二进制文件大小约30%的方法
go build -ldflags "-w -s" xxx.go

交叉编译：指定os和架构
go build编译时可以通过指定GOOS和GOARCH参数来编译指定系统和架构的可执行文件
e.g.:

 GOOS="linux" GOARCH="amd64" go build c1_demo.go

file检测：
c1_demo: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped


用file命令查看可执行文件的类型
e.g.:
file c1_demo
c1_demo: Mach-O 64-bit executable x86_64

*/

func main() {
	// for test golint
    //var Id = 33
    //fmt.Println(Id)
	fmt.Println("Hello, Black Hat Gophers!")
}
