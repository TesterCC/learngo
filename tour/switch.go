package main

import (
	"fmt"
	"runtime"
)

/*
http://127.0.0.1:3999/flowcontrol/9
https://tour.go-zh.org/flowcontrol/9

A switch statement is a shorter way to write a sequence of if - else statements.
It runs the first case whose value is equal to the condition expression.
*/

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
