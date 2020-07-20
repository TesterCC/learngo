package main

import (
	"fmt"
	"time"
)

/*

http://127.0.0.1:3999/methods/19
https://tour.go-zh.org/methods/19

Go 程序使用 error 值来表示错误状态。
error 为 nil 时表示成功；非 nil 的 error 表示失败。

*/

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {

	if err := run(); err != nil {
		fmt.Println(err)
	}
}
