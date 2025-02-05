package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

/*
P19-20 1.1.5 定义参数类型

flag的命令行参数类型是可以自定义的
在Value.Set方法中，只需实现其对应的Value相关的2个接口即可
例子只实现了Value的2个方法。
当然，也可以定制化，无缝接入命令行参数。

cd ~\go_programming_tour\cmd
go run CustomArgs.go -name=TesterCC
*/

type Value interface {
	String() string
	Set(string) error
}

//将原来的字符串变量name修改为类别别名，并为其定义符合Value的两个结构体方法
type Name string

func (i *Name) String() string{
	return fmt.Sprint(*i)
}

func (i *Name) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("name flag already set")
	}
	*i = Name("test custom value: " + value)
	return nil
}

func main() {
	var name Name
	flag.Var(&name, "name", "帮助信息")
	flag.Parse()

	log.Printf("name: %s", name)
}

