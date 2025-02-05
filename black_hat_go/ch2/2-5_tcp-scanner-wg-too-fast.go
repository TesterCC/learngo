package main

import (
	"fmt"
	"net"
	"sync"
)

/*
P27-28  use WaitGroup from the sync package, is a thread-safe way to control concurrency.
issue: 如果多次运行或在不同机器执行可能得到不一致的结果。同时扫描过多的主机或端口可能会导致网络或系统限制，造成结果不正确。

创建WaitGroup后就可以调用它的几种方法了。
第一个是Add(int)，将参数值加到内部的计数器值上。
下一个是Done()，将计数器值减一。最后是Wait()，阻塞所调用的goroutine，直到内部计数器值变为0。
组合这几个函数的调用就能让main goroutine等待所有的goroutine执行完。
*/
const target2 = "scanme.nmap.org"

func main() {
    // WaitGroup是一种结构体，创建后可以在此结构体上调用一些方法
	var wg sync.WaitGroup  // step 1 创建了用作同步计数的sync.WaitGroup

	for i:=1;i<=512;i++ {
		// 按所提供的数字递增内部计数器
		wg.Add(1)   // step 2 每次创建扫描端口的goroutine时，通过wg.Add(1)递增计数器的值
		go func(j int){
			defer wg.Done()   // step 3 且defer语句调用wg.Done()，每当执行完后就使计数器的值递减， 即计数器减1
			address := fmt.Sprintf(target2 +":%d", j)
			fmt.Println(address)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n",j)
		}(i)
	}
	wg.Wait()  // step4 main()函数调用wg.Wait()，会阻止调用它goroutine执行，等待所有的goroutine执行完，在内部计数器的值归0前将不允许进一步执行。
	fmt.Println("[-] Scanner end ...")
}
