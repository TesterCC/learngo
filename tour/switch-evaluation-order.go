package main

import (
	"fmt"
	"time"
)

/*
http://127.0.0.1:3999/flowcontrol/10
https://tour.go-zh.org/flowcontrol/10

switch 的求值顺序
switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
*/

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")

	}
}
