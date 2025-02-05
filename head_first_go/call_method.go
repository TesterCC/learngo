package main

import (
	"fmt"
	"time"
)

// 方法是与特定类型的值关联的函数
func main() {
	var now time.Time = time.Now()
	var year int = now.Year()   // 返回一个代表年的整数
	fmt.Println(year)
}
