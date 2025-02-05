package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 运行在Linux/Unix环境执行（e.g.: CentOS MacOS）
	// tty: pair<type:*os.File, value:"/dev/tty"文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	// r: pair<type: , value: >
	var r io.Reader
	// r: pair<type:*os.File, value:"/dev/tty"文件描述符>

	r = tty

	var w io.Writer
	// r: pair<type:*os.File, value:"/dev/tty"文件描述符>
	w = r.(io.Writer) //  io.Reader 强制转换为 io.Writer

	w.Write([]byte("HELLO, THIS is a TEST !!!\n"))
}
