package main

import "fmt"

/*
https://tour.go-zh.org/flowcontrol/13
https://go.dev/tour/flowcontrol/13

more info:
https://blog.golang.org/defer-panic-and-recover

defer 栈

推迟调用的函数调用会被压入一个栈中。当外层函数返回时，被推迟的调用会按照 后进先出(last-in-first-out order) 的顺序调用。

counting
done
9
8
7
6
5
4
3
2
1
0

*/

func main() {

	fmt.Println("counting") // 1

	for i := 0; i < 10; i++ {
		// 3
		defer fmt.Println(i)
	}

	fmt.Println("done") // 2

}
