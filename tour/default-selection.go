package main

import (
	"fmt"
	"time"
)

/*
http://127.0.0.1:3999/concurrency/6
https://tour.go-zh.org/concurrency/6

默认选择    Default Selection
当 select 中的其它分支都没有准备好时，default 分支就会执行。

为了在尝试发送或者接收时不发生阻塞，可使用 default 分支：

select {
case i := <-c:
    // 使用 i
default:
    // 从 c 中接收会阻塞时执行   receiving from c would block
}
*/

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for  {
		select {
		case <- tick:
			fmt.Println("tick.")
		case <- boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
