package main

import (
	"fmt"
	"strings"
)

// http://govspy.peterbe.com/

func main() {
	// explicitly for every type
	x := 1
	if x != 0 {
		fmt.Println("Yes")
	}

	var y []string // 0
	if len(y) != 0 {
		fmt.Println("this won't printed when len(y) is empty")
	}

	longStr := strings.Repeat("-", 33)
	fmt.Println(longStr)

	fmt.Println(true && false) //  and, false
	fmt.Println(true || false) // or, true
	fmt.Println(!true)         // non, false

}
