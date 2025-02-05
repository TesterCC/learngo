package main

import "fmt"

/*
http://127.0.0.1:3999/flowcontrol/13
https://tour.go-zh.org/flowcontrol/13

more info:
https://blog.golang.org/defer-panic-and-recover

Stacking defers
推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照 后进先出(last-in-first-out order) 的顺序调用。
*/

func main() {

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
