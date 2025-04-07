package main

import (
	"fmt"
	"github.com/pkg/errors"
)

/*
打印错误堆栈信息方法3:
3. 使用 pkg/errors 获取完整堆栈（推荐）
pkg/errors（github.com/pkg/errors）支持错误包装 + 调用栈

优点：
1.包装错误，保留上下文
2.errors.WithStack(err) 自动添加调用栈信息
*/

func doSomething2() error {
	return errors.New("底层错误")
}

func main() {
	err := errors.Wrap(doSomething2(), "doSomething 失败")
	fmt.Printf("%+v\n", err) // 打印完整调用栈
}
