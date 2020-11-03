package main

/*
一个工作需求：
批量执行一个函数列表, 同时最多只能有5个协程同时执行。
by endlex
*/

import (
"fmt"
"sync"
"time"
)

func main() {
	fns := []func(){}
	for i := 0; i < 10; i++ {
		m := i
		fns = append(fns, func() {
			fmt.Println(m)
			time.Sleep(3*time.Second)
		})
	}
	if err := batchFunc(fns, 3); err != nil {
		fmt.Println(err)
	}
}

func batchFunc(fns []func(), limit int) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if e, ok := p.(error); ok {
				err = e
			} else {
				panic(p)
			}
		}
	}()
	if len(fns) < limit {
		limit = len(fns)
	}
	ch := make(chan func())
	go func() {
		defer close(ch)
		for i := range fns {
			ch <- fns[i]
		}
	}()
	g := &sync.WaitGroup{}
	for i := 0; i < limit; i++ {
		g.Add(1)
		go func() {
			for {
				if fn, ok := <-ch; !ok {
					break
				} else {
					fn()
					continue
				}
			}
			g.Done()
		}()
	}
	g.Wait()
	return err
}
