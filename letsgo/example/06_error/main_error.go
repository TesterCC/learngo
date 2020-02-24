package main

import (
	"errors"  // 使用内置的 errors
	"fmt"
)

// Divide compute int a/b
func Divide(a, b int) (int, error) {  // Go一般习惯把最后一个返回值设置为返回错误类型
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil   // 最后一个值返回为nil表示没有错误
}

func main() {
	// fmt.Println(testDefer())
	a, b := 1, 0
	res, err := Divide(a, b)
	if err != nil {  // error不是nil，就说明有错
		fmt.Println(err) // error 类型实现了 Error() 方法可以打印出来
	}
	fmt.Println(res)
}
