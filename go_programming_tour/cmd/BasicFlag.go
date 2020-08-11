package main

import (
	"flag"
	"log"
)

// P11   1.标准库flag的基本使用和长短选项
/*
cd ~/learngo

go run go_programming_tour/cmd/BasicFlag.go -name=TesterCC
go run go_programming_tour/cmd/BasicFlag.go -n=Tester
go run go_programming_tour/cmd/BasicFlag.go -name=TesterCC -n=Tester
go run go_programming_tour/cmd/BasicFlag.go -name CodeCat
go run go_programming_tour/cmd/BasicFlag.go -n CodeCat
*/
func main() {
	var name string

	// 实现对命令行参数name的解析和绑定
	flag.StringVar(&name, "name", "Go Self-learning", "Help Information")
	flag.StringVar(&name, "n", "Go Self-learning Tour", "help information")
	flag.Parse()

	log.Printf("name: %s", name)
}
