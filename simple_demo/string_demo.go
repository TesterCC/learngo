package main

import "fmt"

/*

ref: https://www.cnblogs.com/longzhankunlun/p/13672729.html

[]string 是字符串切片
...string用作参数

...参数语法形成了可变参数的参数。它将接受零个或多个string参数，并将它们作为切片引用
*/

func f(args ...string)  {
	fmt.Println(args)
}

func main() {
	args := []string{"a","b"}
	f(args...)
	args2 := []string{"a","b","c","d"}
	f(args2...)
	var args0 []string
	f(args0...)
}