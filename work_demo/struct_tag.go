package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// ref: https://zhuanlan.zhihu.com/p/258978922
// 因此 encoding/json 在将对象转化 json 字符串时，只要发现对象里的 Addr 为 false， 0， 空指针，空接口，空数组，空切片，空映射，空字符串中的一种，就会被忽略。
// omitempty
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

func main() {
	p1 := Person{
		Name: "Bob",
		Age:  22,
	}

	// 序列化： obj ->  字节序列
	data1, err := json.Marshal(p1) // 类似python, json.dumps()

	if err != nil {
		panic(err)
	}

	// p1 没有 Addr，就不会打印了
	fmt.Printf("%s\n", data1)

	p2 := Person{
		Name: "Alice",
		Age:  18,
		Addr: "Zoo, China", // 需要加 , 号结尾
	}

	// json序列化： obj -> 字节序列
	data2, err := json.Marshal(p2) // 类似python, json.dumps()

	if err != nil {
		panic(err)
	}

	// p1 没有 Addr，就不会打印了
	fmt.Printf("%s\n", data2)

	// 从结构体中取出tag
	p := reflect.TypeOf(Person{})

	name, _ := p.FieldByName("Name")
	age, _ := p.FieldByName("Age")
	addr, _ := p.FieldByName("Addr")

	// %q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	// 利用反射从结构体中取出Tag
	fmt.Printf("%q\n", name.Tag)
	fmt.Printf("%q\n", age.Tag)
	fmt.Printf("%q\n", addr.Tag)
}
