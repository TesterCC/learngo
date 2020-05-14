package main

import "fmt"

type People interface {
}

// P15 special syntax
func foo(i interface{}) {
	// retrieve
	switch v := i.(type) {
	case int:
		fmt.Println("I'm an integer!")
	case string:
		fmt.Println("I'm a string!")
	default:
		fmt.Println("Unknown type!")
	}
}

func main() {
	// need to learn go interface to understand
}
