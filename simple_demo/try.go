package main

import "fmt"

// Endlex说的坑

func main() {
	a := []int{1,2,3}
	b := a
	c := a
	//fmt.Println(a,b,c)
	b = append(b, 1)
	c[1] = 5
	fmt.Println(a)   // [1 5 3]
	fmt.Println(b)   // [1 2 3 1]  另外copy了一份
	fmt.Println(c)   // [1 5 3]
}
