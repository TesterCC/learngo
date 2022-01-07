package main

import (
	"fmt"
	"reflect"
)

// 23-golang反射reflect机制用法
type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called ...")
	fmt.Printf("%v\n", this)
}

func main() {
	// 定义一个User，通过反射机制动态把对象解析出来
	user := User{1, "Alice", 18} // https://www.bilibili.com/video/BV1gf4y1r79E?p=23 5:18
	DoFieldAndMethod(user)
}

// 定义一个方法用于传递任何数据类型
func DoFieldAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is : ", inputType)

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is : ", inputValue)

	// 通过type获取里面的字段
	// 1. 获取interface的reflect.Type, 通过Type得到NumField，进行遍历
	// 2. 得到每个field和数据类型
	// 3. 通过field有一个Interface()方法得到对应的value

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n",  field.Name, field.Type, value)
	}

	// 通过type获取里面的方法，调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
