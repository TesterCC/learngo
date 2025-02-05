package main

import (
	"errors"
	"fmt"
)

//P17-18  1.2.4 错误处理，Go没有try/catch/finally错误处理语法。

// Go使用以下接口声明定义内置错误类型
//type error interface {
//	Error() string
//}

// a custom error you can use, a user-defined string type
// 创建一个名为MyError的用户定义字符串并为该类型实现方法Error()
type MyError string

func (e MyError) Error() string {
	return string(e)
}

func foo2() error {
	return errors.New("Some Error Occurred")
}

func main() {
	if err := foo2(); err != nil {
		//Handle the error
		fmt.Println(err) // print error info
	}
}
