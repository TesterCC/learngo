package main

import (
	"fmt"
	"sync"
)

/*
一个工作需求：
批量执行一个函数列表, 同时最多只能有5个协程同时执行。
by pencil
*/

func RunParallel(inputChan chan interface{}, worker func(input interface{}) error, concurrency int) {
	wg := sync.WaitGroup{}
	for i := 1; i <= concurrency; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("worker-%d start\n", i)
			for {
				select {
				case v, ok := <-inputChan:
					if !ok {
						fmt.Println("input channel closed")
						return
					}
					err := worker(v)
					if err != nil {
						fmt.Println(err)
						continue
					}
				}
			}
		}(i)
	}
	wg.Wait()
}

func main() {
	inputs := []int{1, 2, 3, 4}
	double := func(input interface{}) error {
		v := input.(int)
		result := v * 2
		fmt.Printf("input: %v, result: %v\n", input, result)
		return nil
	}
	inputChan := make(chan interface{})

	go func() {
		defer close(inputChan)
		for _, input := range inputs {
			inputChan <- input
		}
	}()
	RunParallel(inputChan, double, 3)
	fmt.Println("done")
}
