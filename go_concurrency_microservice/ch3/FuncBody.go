package main

import "fmt"

// P62  Go语言中所有函数都可以实现接口。
// 实践一个通过函数实现接口的例子。

type Printer interface {
	// 打印方法，将输入的参数直接打印到命令行中
	Print(interface{})
}

// 由于 函数不能直接实现接口，需要先 将函数定义为类型，再 使用定义好的函数类型实现接口
// 函数定义为类型
type FuncCaller func(p interface{})

// 实现Printer的Print方法
func (funcCaller FuncCaller) Print(p interface{}) {
	// 调用funcCaller函数本体
	funcCaller(p)
}
// 接口的实现由直接调用定义好的函数类型来完成

func main() {
	// 要先定义好FuncCaller函数类型用于提供具体逻辑处理，然后就可以将任意匹配函数强转为FuncCaller
	// 如果通过一个匿名函数实现FuncCaller
	var printer Printer
	// 将匿名函数强转为FuncCaller，赋值给printer
	printer = FuncCaller(func(p interface{}) {
		fmt.Println(p)
	})
	printer.Print("Golang is Good!")
}

// 通过函数类型实现接口，可以将函数作为接口来使用，在运行时可以通过替换具体的实现函数，实现类似多太的效果。
