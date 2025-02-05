package main

import (
	"fmt"
	"os"
)

// gopl P3 output command args
/*
➜  ch1 git:(master) ✗ go build echo1.go
➜  ch1 git:(master) ✗ ./echo1 1 2 3

*/

// echo1 输出其命令行参数
func main() {
	// 隐式初始化为空字符串""
	var s, sep string
	// 命令行参数以os包中Args名字的变量供程序访问，在os包外，使用os.Args这个名字, os.Args是一个字符串slice
	// go的os.Args类似python中的sys.argv
	// for是go里的唯一循环语句
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		//fmt.Println(s,sep)
	}
	fmt.Println(s)
}
