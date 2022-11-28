package main

import (
	"fmt"
	"time"
)

// P15 use goroutines 可使用goroutine并发执行代码。goroutine通常被称为轻量级线程（比创建实际线程成本低）
// 通过在被调用的方法或函数之前使用go关键字创建goroutine实现并发。

func f(){
	fmt.Println("f function")
}

func main() {
	go f()  // use the go keyword before the call to method or function you wish
	// 如果不暂停，则该程序很可能在函数f()执行完之前就推出，并永远不会看到其结果显示到标准输出。
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}

