package main

import "fmt"

type User struct {
	Id int
	Name string
	Age int
}

func (this *User) Call() {
	fmt.Println("user is called ...")
	fmt.Printf("%v\n", this)
}

func main() {
	// 定义一个User，通过反射机制动态把对象解析出来
	user := User{1, "Alice", 18}   // https://www.bilibili.com/video/BV1gf4y1r79E?p=23 5:18

}

// 定义一个方法用于传递任何数据类型
func DoFiledAndMethod(input interface{})  {
	// 获取input的type

	// 获取input的value

	// 通过type获取里面的字段

	// 通过type获取里面的方法，调用

}