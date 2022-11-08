package main

import (
	"fmt"
	"runtime"
)

/*
http://www.nfangbian.com/go/bf/runtime.html

use runtime get basic info

runtime.GOOS 获取操作系统类型，如：darwin、freebsd、linux
runtime.GOARCH 获取操作系统内核，如：amd64、386、arm、s390x
*/

func main() {

	// get OS
	os := runtime.GOOS
	//fmt.Println(os)
	fmt.Printf("Operation System: %s\n", os)

	arch := runtime.GOARCH
	fmt.Printf("System ARCH: %s\n", arch)
	//fmt.Println(arch)

	// get cpu count
	cpuNum := runtime.NumCPU()
	//fmt.Println(cpuNum)
	fmt.Printf("CPU Num: %d\n", cpuNum)

	// get GOROOT dir
	goRoot := runtime.GOROOT()
	//fmt.Println(goRoot)
	fmt.Printf("GOROOT dir: %s\n", goRoot)

	// get go version
	goVersion := runtime.Version()
	//fmt.Println(goVersion)
	fmt.Printf("Go Version: %s\n", goVersion)

	// get launch goroutine count
	goroutineCount := runtime.NumGoroutine()
	//fmt.Println(goroutineCount)
	fmt.Printf("Goroutine num: %d\n", goroutineCount)

	// get called cgo count
	cgoCount := runtime.NumCgoCall()
	//fmt.Println(cgoCount)
	fmt.Printf("cgo num: %d\n", cgoCount)

}
