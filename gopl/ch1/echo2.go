package main

import (
	"fmt"
	"os"
)

// P5 echo2 输出其命令行参数

func main() {
	s, sep := "",""
	// 每次迭代，range产生一对值：索引和这个索引处元素的值。
	// 因为go不允许存在无用的临时变量，所以用空标识符"_"
	for _,arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
