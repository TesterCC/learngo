package main

import (
	"fmt"
	"time"
)

/*
http://127.0.0.1:3999/concurrency/1
https://tour.go-zh.org/concurrency/1

A goroutine is a lightweight thread managed by the Go runtime.
go f(x, y, z) 会启动一个新的 Go 程并执行 f(x, y, z)

f, x, y 和 z 的求值发生在当前的 Go 程中，而 f 的执行发生在新的 Go 程中。

goroutine 在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。
sync 包提供了这种能力，不过在 Go 中并不经常用到，因为还有其它的办法（见下一页）。
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")  // starts a new goroutine running
	say("hello")

}
