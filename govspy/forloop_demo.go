package main

import (
	"fmt"
	"strings"
)

// Go has only one type of loop and that's the for loop.

func main() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i += 1
	}

	longStr := strings.Repeat("-", 33)
	fmt.Println(longStr)

	// same thing more but more convenient
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

}
