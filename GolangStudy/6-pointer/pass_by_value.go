package main

import "fmt"

/*
ref: https://www.bilibili.com/video/BV1gf4y1r79E?p=10
*/


func changeValue2(p int) {
	p = 10
}

func main() {
	var a int = 1
	changeValue2(a)
	fmt.Println("a = ", a)   // a = 1, 值传递，再怎么修改p也不影响a的值
}

