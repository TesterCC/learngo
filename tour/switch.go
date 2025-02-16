package main

import (
	"fmt"
	"runtime" // 提供与 Go 程序运行时环境相关的功能，如获取系统信息。
)

/*
https://tour.go-zh.org/flowcontrol/9
https://go.dev/tour/flowcontrol/9

switch 分支
switch 语句是编写一连串 if - else 语句的简便方法。它运行第一个 case 值 值等于条件表达式的子句。

Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只会运行选定的 case，而非之后所有的 case。
在效果上，Go 的做法相当于这些语言中为每个 case 后面自动添加了所需的 break 语句。
在 Go 中，除非以 fallthrough 语句结束，否则分支会自动终止。
Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不限于整数。

A switch statement is a shorter way to write a sequence of if - else statements.
It runs the first case whose value is equal to the condition expression.
*/

func main() {
	fmt.Println("System Architecture:", runtime.GOARCH) // mac silicon is arm64

	fmt.Print("Go runs on system: ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
