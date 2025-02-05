package main

import "fmt"

// P88-89


func setVTo1(v *int){
	*v = 1
}


func setVTo2(v *int){
	*v = 2
}

func main(){
	v := new(int)
	go setVTo1(v)
	go setVTo2(v)
	fmt.Println(*v)   // *v 输出可能是0、1、2，但大多数情况是0.
}