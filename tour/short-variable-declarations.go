package main

import "fmt"

func main() {
	var i, j int = 1,2
	k := 3   // Attention: Outside a function, := construct is not available.
	c, python, java := true, false, "no!"

	fmt.Println(i,j,k,c,python,java)
}
