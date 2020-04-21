package main

import (
	"fmt"
	"time"
)

// ref:https://www.bilibili.com/video/BV1gz41187kG

func goFunc(i int){
	fmt.Println("goroutine", i, "...")
}

func main() {
	for i := 0; i < 10000; i++ {
		go goFunc(i)   //通过go关键字，开启一个并发协程
	}

	time.Sleep(time.Second)
}
