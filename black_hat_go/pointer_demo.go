package main

import "fmt"

/*
P12
*/
func main() {
	var count = int(42)
	ptr := &count // create a pointer
	fmt.Println(*ptr)
	*ptr = 100 // assign a new value to the memory
	fmt.Println(count)
}
