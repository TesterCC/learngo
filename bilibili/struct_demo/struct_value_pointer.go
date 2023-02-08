package main

import "fmt"

/*
Go struct 该用值定义方法还是用指针
ref: https://www.bilibili.com/video/BV1o24y167ac

接收器该定义为值还是指针？

1.如果方法需要修改接收器，则接收器使用指针；
2.从性能考量：
 a. 如果结构体很大，应该使用指针接收器，会减少内存拷贝
 b. 对于基本类型、切片和小结构等类型，使用值接收器是非常廉价的
3.方法集中的接收器应该保持一致，要么全为值，要么全为指针
*/

type Person struct {
	Name string
}

// method 1  Go语言函数传参是值传递
func (p Person) SetName1(name string) {
	// p is receiver，形参p是实参person的副本，此方法修改内部的形参p并不改变外部的实参person
	p.Name = name
}

// method 2
func (p *Person) SetName2(name string) {
	// 传入的是指针，实参&person，虽与形参p不是一个变量，但都指向person
	p.Name = name
}

func main() {
	person := Person{}
	person.SetName1("天选打工人")
	fmt.Println("SetName1: ", person.Name)
	person.SetName2("地选打工人")
	fmt.Println("SetName2: ", person.Name)
}
