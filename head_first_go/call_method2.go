package main

import (
	"fmt"
	"strings"
)

// 字符串替换，类似python的replace("a","b")
func main() {

	broken := "G# r#coks!"
	replacer := strings.NewReplacer("#","o")  // return a strings.Replacer
	fixed := replacer.Replace(broken)   // call strings.Replacer Replace() method
	fmt.Println(fixed)
}
