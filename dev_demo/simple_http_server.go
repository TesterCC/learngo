package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// ref: https://www.cnblogs.com/Yellow0-0River/p/7543869.html
// 在浏览器里输入http://127.0.0.1:8080即可浏览文件，这些文件正是当前目录在HTTP服务器上的映射目录。
// 复杂功能参考：https://github.com/projectdiscovery/simplehttpserver

func getCurrentPath() string {
	s, _ := exec.LookPath(os.Args[0])
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func main() {
	path := getCurrentPath()
	fmt.Println("Current Path:" + path)
	// 使用 http.FileServer 文件服务器将当前目录作为根目录（/目录）的处理器，访问根目录，就会进入当前目录
	//http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// GOOS="linux" GOARCH="amd64" go build simple_http_server.go

// todo update to commmand tool