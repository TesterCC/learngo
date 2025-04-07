package main

import (
	"errors"
	"fmt"
)

/*
打印错误堆栈信息方法1:
1. fmt.Errorf + errors.Unwrap()
fmt.Errorf 可以添加上下文信息
*/

func doSomething() error {
	return errors.New("底层错误")
}

func main() {
	// 可增强上下文信息，但不包含调用栈
	err := fmt.Errorf("doSomething 失败: %w", doSomething()) // %w 让 errors.Unwrap(err) 取出底层错误
	fmt.Println(err)                                       // 包含上下文信息
}
