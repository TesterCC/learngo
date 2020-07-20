package main

import "fmt"

// P23  for 遍历数组的3种方式

func main() {
	var a = [...]int{1, 2, 3}
	b := a
	c := a

	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}

	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}

	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]: %d\n", i, c[i])
	}
}
