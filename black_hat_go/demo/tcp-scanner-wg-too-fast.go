package main

import (
	"fmt"
	"net"
	"sync"
)

/*
P27-28  use WaitGroup from the sync package, is a thread-safe way to control concurrency.
*/
func main() {
	var target = "testphp.vulnweb.com"

	var wg sync.WaitGroup  // step 1

	for i:=1;i<=1024;i++ {
		wg.Add(1)   // step 2
		go func(j int){
			defer wg.Done()   // step 3
			address := fmt.Sprintf(target+":%d", i)
			fmt.Println(address)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n",j)
		}(i)
	}
	wg.Wait()  // step4
}
