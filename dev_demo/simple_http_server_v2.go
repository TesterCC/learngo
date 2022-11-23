package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Options struct {
	ListenAddress string
	Version       bool
}

// Version is the current version
const Version = `0.2.20221118`

func showBanner() {
	fmt.Printf("Custom Simple HTTP Server, version: %s.\n", Version)
	fmt.Printf("Developers assume no liability and are not responsible for any misuse or damage.\n")
}

// ParseOptions parses the command line options for application
func ParseOptions() *Options {
	options := &Options{}
	flag.StringVar(&options.ListenAddress, "listen", "0.0.0.0:8080", "Address:Port")
	flag.BoolVar(&options.Version, "version", false, "Show version")
	flag.Parse()

	showBanner()

	if options.Version {
		fmt.Printf("Current Version: %s.\n\n", Version)
		fmt.Printf("Please use -h to get usage.\n")
		os.Exit(0)
	}

	return options
}

func main() {

	options := ParseOptions()
	fmt.Printf("[+] listen server at %s\n", options.ListenAddress)

	// can delete
	ret := strings.Split(options.ListenAddress,":")
    listenPort := ret[len(ret)-1]
	fmt.Println("[D] listen port:", listenPort)

	currentPath := "."
	if path, err := os.Getwd(); err == nil {
		currentPath = path
	}
	fmt.Println("[+] current path: " + currentPath)

	// 使用 http.FileServer 文件服务器将当前目录作为根目录（/目录）的处理器，访问根目录，就会进入当前目录
	//http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/", http.FileServer(http.Dir(".")))
	//err := http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(options.ListenAddress, nil)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}
// go cross compiling
// GOOS="linux" GOARCH="amd64" go build simple_http_server.go
//默认go build包含debug信息和符号表，减少生成二进制文件大小约30%的方法
// go build -ldflags "-w -s" xxx.go

// 交叉编译+减少文件大小
// GOOS="linux" GOARCH="amd64" go build -o shs -ldflags "-w -s" simple_http_server_v2.go

