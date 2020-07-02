package main

import "fmt"

/*
go 里区分对待异常(panic)和错误(error)的，绝大部分场景下我们使用的都是错误，只有少数场景下发生了严重错误我们想让整个进程都退出了才会使用异常。

对于一般不太严重的场景，返回错误值 error 类型 (业务绝大部分场景)
对于严重的错误需要整个进程退出的场景，使用 panic 来抛异常，及早发现错误
如果希望捕获 panic 异常，可以使用 recover 函数捕获，并且包装成一个错误返回
web 框架等会帮你捕获 panic 异常，然后返回客户端一个 http 500 状态码错误
*/


func MustDivide2(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}

// 使用 recover 可以把一个 panic 包装成为 error 再返回，而不是让进程退出。
func Divide3(a, b int) (res int, e error) {
	defer func() {
		// 捕获了 panic 异常并且返回了一个错误，代码也可以正常执行而不会退出
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()
	res = MustDivide2(a, b)   // 遇到严重错误会panic
	return // 命名返回值不用加上返回的参数
}

func main() {
	Divide3(1,0)
}

