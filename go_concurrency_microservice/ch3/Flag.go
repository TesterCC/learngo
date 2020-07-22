package main

import (
	"flag"
	"fmt"
)

/*
P47 使用flag从命令行中读取

不指定参数则传入默认值

./Flag -surname="王" -personalName="二小" -id=3
./Flag --surname="王" --personalName="二小" --id=4

*/

func main() {
	// 定义一个类型为string,名称为surname（姓式）的命令行参数
	// 参数依次是命令行参数的名称，默认值，提示
	surname := flag.String("surname", "李", "您的姓")

	// 定义一个类型为string,名称为surname（姓式）的命令行参数
	// 除了返回指针类型结果，还可以直接传入变量地址获取参数值
	var personalName string
	flag.StringVar(&personalName, "personalName", "小明","您的名")

	// 定义一个类型为int，名称为id的命令行参数
	id := flag.Int("id", 0, "您的ID")

	// 解析命令行参数
	flag.Parse()
	fmt.Printf("I am %v %v, and my id is %v\n", *surname, personalName, *id)
}
