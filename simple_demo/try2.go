package main

import "fmt"

func main() {
	a := [...]int{1,2,3}
	a1 := a[0:2]
	a2 := a[1:3]

	fmt.Printf("a:%p; a1:%p; a2:%p\n", &a, a1, a2)
	fmt.Printf("a[1]: %p\n", &a[1])
	fmt.Printf("cap(a1): %d , cap(a2):%d\n", cap(a1),cap(a2))
}
