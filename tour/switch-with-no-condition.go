package main

import (
	"fmt"
	"time"
)

/*
http://127.0.0.1:3999/flowcontrol/11
https://tour.go-zh.org/flowcontrol/11

没有条件的 switch 同 switch true 一样。

这种形式能将一长串 if-then-else 写得更加清晰
*/

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	
}
