package main

import "fmt"

// P43

func getName() (string, string) {
	return "Alice", "Albert"
}

func main() {
	surname, _ := getName()
	_, personalname := getName()

	fmt.Printf("My surname name is %v and my personal name is %v.\n", surname, personalname)
}
