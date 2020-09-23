package main

import (
	"flag"
	"log"
)

// P11   1.标准库flag的基本使用和长短选项

/*
命令行参数支持的3种标志语法
-flag : just support boolean
-flag x : just support non-boolean
-flag=x : all support

标准库flag提供了多种类型参数的绑定方式

cd ~/learngo

go run go_programming_tour/cmd/BasicFlag.go -name=TesterCC
go run go_programming_tour/cmd/BasicFlag.go -n=Tester
go run go_programming_tour/cmd/BasicFlag.go -name=TesterCC -n=Tester
go run go_programming_tour/cmd/BasicFlag.go -name CodeCat
go run go_programming_tour/cmd/BasicFlag.go -n CodeCat
*/
func main() {
	var name string

	// 调用StringVar方法，实现对命令行参数name的解析和绑定
	// 形参含义，分别为命令行标识位的 名称name、默认值value、帮助信息usage
	flag.StringVar(&name, "name", "Go Self-learning", "Help Information")
	flag.StringVar(&name, "n", "Go Self-learning Tour", "help information")
	flag.Parse()

	log.Printf("name: %s", name)
}
