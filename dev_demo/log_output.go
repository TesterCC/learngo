package main

import (
	"io"
	"log"
	"os"
)

/*
将日志同时输出到控制台和文件(常见开发问题)

log包可以通过SetOutput()方法指定日志输出的方式（Writer），但是只能指定一个输出的方式（Writer）。我们利用io.MultiWriter()将多个Writer拼成一个Writer使用的特性，把log.Println()输出的内容分流到控制台和文件当中。

 */
func main() {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.Println("log test")
}
