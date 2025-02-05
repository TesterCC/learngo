package main

import (
	"fmt"
	"time"
)

func goFunc(i int)  {
	fmt.Println("goroutine", i, "...")
}

func main() {
	for i := 0; i < 10000; i++ {
		go goFunc(i)  // 开启一个并发协程   这是一个单独的执行流程体
	}

	time.Sleep(time.Second)
}
