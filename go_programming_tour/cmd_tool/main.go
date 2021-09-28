package main

import (
	"learngo/go_programming_tour/cmd_tool/cmd"
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

注意：cmd_tool 目录下的main.go才是命令行集合工具的入口

Testcase:
go run main.go help word
go run main.go word -s=TesterCC -m=1
go run main.go word -s=TesterCC -m=2
go run main.go word -s=tester_cc -m=3
go run main.go word -s=tester_cc -m=4
go run main.go word -s=TesterCC -m=5

go run main.go help time
go run main.go time now
go run main.go time calc -c="2029-09-04 12:02:33" -d=5m
go run main.go time calc -c="2029-09-04 12:02:33" -d=-2h

实际上在使用标准库 time 时是存在遇到时区问题的风险的

不同的国家（有时甚至是同一个国家内的不同地区）使用着不同的时区。对于要输入和输出时间的程序来说，必须对系统所处的时区加以考虑。而在 Go 语言中使用 Location 来表示地区相关的时区，一个 Location 可能表示多个时区。

在标准库 time 上，提供了 Location 的两个实例：Local 和 UTC。Local 代表当前系统本地时区；UTC 代表通用协调时间，也就是零时区。

在默认值上，标准库 time 使用的是 UTC 时区，即0时区。

Local 是如何表示本地时区的

时区信息既浩繁又多变，Unix 系统以标准格式存于文件中，这些文件位于 /usr/share/zoneinfo，而本地时区可以通过 /etc/localtime 获取，这是一个符号链接，指向 /usr/share/zoneinfo 中某一个时区。
比如我本地电脑指向的是：/var/db/timezone/zoneinfo/Asia/Shanghai。
*/

