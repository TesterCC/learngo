package main

import (
	"fmt"
	"reflect"
)

// ref: https://zhuanlan.zhihu.com/p/258978922
// 因此 encoding/json 在将对象转化 json 字符串时，只要发现对象里的 Addr 为 false， 0， 空指针，空接口，空数组，空切片，空映射，空字符串中的一种，就会被忽略。

// 实战一下：利用 Tag 搞点事情

// 改造下 Person 结构体为 Person2，给每个字段加上 tag 标签，三个字段的tag 都有 label 属性，而 Gender 多了一个 default 属性，意在指定默认值。
type Person2 struct {
	Name string `label:"Name is: "`
	Age  int    `label:"Age is: "`
	Addr string `label:"Gender is: " default:"unknown" `
}

// Print function  在实际开发中可用
func Print(obj interface{}) error {
	// get Value
	v := reflect.ValueOf(obj)

	// parse field
	for i := 0; i < v.NumField(); i++ {
		// get tag
		field := v.Type().Field(i)
		tag := field.Tag

		// parse label and default
		label := tag.Get("label")
		defaultValue := tag.Get("default")

		value := fmt.Sprintf("%v", v.Field(i))

		if value == "" {
			// 如果没有指定值，则用默认值替代
			value = defaultValue
		}

		fmt.Println(label + value)

	}
	return nil

}

func main() {
	p2 := Person2{
		Name: "Alice",
		Age:  18,
	}
	err := Print(p2)
	if err != nil {
		return 
	}
}