package main

import (
	"flag"
	"fmt"
	"log"
)

// P12  2.子命令的使用
/*
一个工具可能包含大量相关联的功能命令，以形成工具集

test：
cd ~/learngo
go run go_programming_tour/cmd/SubCmd.go go -name=CC
go run go_programming_tour/cmd/SubCmd.go php -n=CCC

go run go_programming_tour/cmd/SubCmd.go go -n=Code
go run go_programming_tour/cmd/SubCmd.go python -n=CodeCat
*/

var name string

func main() {
	flag.Parse()  // 将命令行解析为定义的标志，以便后续参数使用

	// 该方法会返回带有这个指定名称和错误处理属性的空命令集，相当于创建一个新的命令集去支持子命令，第二个参数ErrorHandling用于指定处理异常错误
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "Go语言","Help Information")
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "name", "PHP语言","Help Information")
	pythonCmd := flag.NewFlagSet("python", flag.ExitOnError)
	pythonCmd.StringVar(&name, "name", "Python语言","Help Information")
    // 因为 ErrorHandling 传递的是 ExitOnError 级别的命令，因此当识别出传递的命令行参数标志是未定义的时，会直接退出程序并提示错误。

	args := flag.Args()
	fmt.Println(args)
	fmt.Println(args[0], args[1:])
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "php":
		_ = phpCmd.Parse(args[1:])
	case "python":
		_ = pythonCmd.Parse(args[1:])
	default:
		fmt.Println("There is no args value")
	}

	log.Printf("name: %s", name)
}
