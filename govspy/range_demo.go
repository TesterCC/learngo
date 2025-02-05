package main

import "fmt"

func main() {
	names := []string{
		"Alice",
		"Bob",
		"David",
		"Eric",
		"Jack",
		"Peter",
	}

	for i, name := range names {
		fmt.Printf("%d. %s\n", i, name)
	}
}
