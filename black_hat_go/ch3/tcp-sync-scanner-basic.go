package main

import (
	"fmt"
	"sync"
)

/*
P28-29 Port Scanning Using a Worker Pool
*/

// A worker function for processing work
// channel will be used to receiver work
// WaitGroup will be used to track when a single work item has been completed.
func worker(ports chan int, wg *sync.WaitGroup){
	for p := range ports{
		fmt.Println(p)    // print port
		wg.Done()
	}
}

// manage the workload and provide work to your worker(int, *sync.WaitGroup) function.
func main() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 0; i< cap(ports); i++ {
		go worker(ports, &wg)
	}

	for i:=1; i<=1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}
