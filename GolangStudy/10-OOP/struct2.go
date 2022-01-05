package main

import "fmt"

/*
17 - struct基本定义与使用
ref: https://www.bilibili.com/video/BV1gf4y1r79E?p=17
*/

// 定义要给结构体
type Book struct {
	title string
	auth string
}

func changeBook(book Book)  {
	// 传递一个book的副本  并不能改变book1
	book.auth = "test"
}

func changeBook2(book *Book)  {
	// pass by pointer 指针传递
	book.auth = "OpenSource"

}

func main() {
	var book1 Book
	book1.title = "Golang"
	book1.auth = "Google"

	fmt.Printf("%v\n, %T\n", book1,book1)
	changeBook(book1)
	fmt.Printf("%v\n", book1)
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}

