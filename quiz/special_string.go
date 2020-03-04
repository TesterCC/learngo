package main

import "fmt"

func main() {

	a := 0
	fmt.Println(string(a))   // 实际输出空字符串，IDE输出结果为乱码, 换成 ascii 就存的 “\u0000”
	
}
