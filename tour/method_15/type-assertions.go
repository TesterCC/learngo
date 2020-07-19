package main

import "fmt"

/*

http://127.0.0.1:3999/methods/15
https://tour.go-zh.org/methods/15

类型断言 提供了访问接口值底层具体值的方式。

为了 判断 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。
*/

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s,ok := i.(string)
	fmt.Println(s,ok)

	f, ok := i.(float64)
	fmt.Println(f,ok)

	f = i.(float64)  // call panic
	fmt.Println(f)
}
