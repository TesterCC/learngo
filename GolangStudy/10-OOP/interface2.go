package main

import "fmt"

/*
21-interface空接口万能类型与类型断言机制

ref: bilibili.com/video/BV1gf4y1r79E?p=21

通用万能类型 interface{} 空接口
实现了interface{} 就可以用interface{}类型引用任意的数据类型
 */

// interface{} 是万能数据类型
func myFunc(arg interface{})  {
	fmt.Println("myFunc is called ...")
	fmt.Println(arg)

	// interface{} 该如何区分 此时引用的底层数据类型到底是什么？ 3：56

}

type Book2 struct {
	auth string
}

func main() {
	book := Book2{"Golang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
