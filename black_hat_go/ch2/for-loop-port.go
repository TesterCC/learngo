package main

import (
	"fmt"
	"time"
)

// P25
var target = "172.16.12.10"
func main() {
	for i:= 1; i< 10; i++ {
		address := fmt.Sprintf(target+":%d", i)  // 变量拼接
		fmt.Println(address)
		//time.Sleep(2 * time.Second)   // 增加延时
		time.Sleep(1500 * time.Millisecond)
	}
}
