package main

import "fmt"

/*
ref: https://www.bilibili.com/video/BV1gf4y1r79E?p=10
*/

// 引用传递 / 指针传递
// p是指向int型的指针类型
func changeValue(p *int) {
	*p = 10
}

func main() {
   var a int = 1
   changeValue(&a)   // &表示传址，传的a的地址
   fmt.Println("a = ", a)    // a = 10
}


