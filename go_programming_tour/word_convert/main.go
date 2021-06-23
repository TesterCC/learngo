package main

import (
	"learngo/go_programming_tour/word_convert/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Excute err: %v", err)
	}
}

/*
基于第三方开源库 Cobra 和标准库 strings、unicode 实现了多种模式的单词转换，非常简单，也是在日常的工作中较实用的一环，因为我们经常会需要对输入、输出数据进行各类型的转换和拼装。
Testcase:
go run main.go help word
go run main.go word -s=TesterCC -m=1
go run main.go word -s=TesterCC -m=2
go run main.go word -s=tester_cc -m=3
go run main.go word -s=tester_cc -m=4
go run main.go word -s=TesterCC -m=5
*/
