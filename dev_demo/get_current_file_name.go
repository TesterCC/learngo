package main

import (
	"fmt"
	"path"
	"runtime"
)


func main() {

	// 当前运行文件全路径和名称
	_, fullFilename, _, _ := runtime.Caller(0)
	fmt.Println("full file name:",fullFilename)

	// 当前运行文件名
	var filename string
	filename = path.Base(fullFilename)
	fmt.Println("current filename:",filename)

}
