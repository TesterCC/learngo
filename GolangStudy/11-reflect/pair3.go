package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book struct {
	
}

func (this *Book) ReadBook()  {
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	// b: pair<type:Book, value:Book{}地址>
	b := &Book{}

	// r: pair<type: , value: >
	var r Reader
	// r: pair<type:Book, value:Book{}地址>
	r = b
	r.ReadBook()

	var w Writer
	// w: pair<type:Book, value:Book{}地址>
	w = r.(Writer)    // 此处的断言(强制类型转换)为什么会成功？  因为w和r具体的type是一致的 Book

	w.WriteBook()
}
