package main

import "fmt"

/*
22-变量的内置pair结构详细说明
ref:
https://www.bilibili.com/video/BV1gf4y1r79E?p=22

变量的pair包含：type 和 value

type: static type , concrete type
*/
func main() {
	var a string

	// pair<static type:string, value: "security">
	a = "security"

	// pair<type:string, value: "security">
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
