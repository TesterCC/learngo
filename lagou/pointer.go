package main

import "fmt"

func main() {
	name := "TesterCC"
	p := &name
	fmt.Println("name is", *p)
	*p = "CodeCat"
	fmt.Println("name is", name)
}
