package main

import (
	"fmt"
	"io/ioutil"
)

/*
if的条件里可以赋值
if的条件里赋值的变量变量作用域就在这个if语句里
*/


func main() {
	const filename = "abc.txt"   // terminal can use relative path   // IDE need abs path
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println("cannot print file contents:", err)
	} else {
		fmt.Printf("%s\n", string(contents))   // contents byte
	}
}