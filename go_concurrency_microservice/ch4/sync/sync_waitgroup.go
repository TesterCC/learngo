package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*
P97 WaitGroup并发等待组

使用sync.WaitGroup的goroutine会等待预设好数量的goroutine都提交执行结束后，才会继续往下执行代码

sync.WaitGroup 适用于执行批量操作，等待所有goroutine执行结束后统一返回结果

4-7 sync.WaitGroup 阻塞主goroutine直到其它goroutine执行结束
*/

func main() {
	var waitGroup sync.WaitGroup

	// add and wait goroutine count increase to 5
	waitGroup.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println("work " + strconv.Itoa(i) + " is done at " + time.Now().String())
			// wait 1 second and count -1
			time.Sleep(time.Second)
			waitGroup.Done()  // 需要等待waitGroup中的等待数变为0之后才能继续往后执行
		}(i)
	}

	waitGroup.Wait()
	fmt.Println("all works are done at " + time.Now().String())
}
