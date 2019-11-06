package main

import "fmt"

//var i,j int = 1,2 // If an initializer is present, the type can be omitted
var i,j = 1,2

func main() {
	var c,python,java = true, false, "No!"
	fmt.Println(i,j, c, python, java)
}
