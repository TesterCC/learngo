package main

import "fmt"

// P16-17 use channels
func strlen(s string, c chan int){
	c <- len(s)
}

func main() {
	c := make(chan int)
	go strlen("Salutations", c)
	go strlen("World", c)
	x,y := <-c, <-c
	fmt.Println(x,y,x+y)
}
