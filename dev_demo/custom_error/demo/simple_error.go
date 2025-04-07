package main

import (
	"fmt"
)

// 在 Go 中，error 是一个接口，默认不包含堆栈跟踪，只提供简单的错误信息。

// 自定义错误类型
type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func doSomething(flag bool) error {
	if !flag {
		return &MyError{"发生了一个自定义预期错误"}
	}
	return nil
}

func main() {
	err := doSomething(false)
	if err != nil {
		fmt.Println("错误:", err)
	}
}
