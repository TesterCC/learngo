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

// 与其它语言不通，Go的内置错误类型没有隐式包含堆栈跟踪以帮助查明错误的上下文或位置。
// 尽管可以生成一个错误处理方法并在应用中把它赋给一个自定义类型，但其实现方式需要开发人员去摸索。
