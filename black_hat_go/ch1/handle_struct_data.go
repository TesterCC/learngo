package main

import (
	"encoding/json"
	"fmt"
)

//P19 可以将字符串转换为结构体，也可以将结构体转换为字符串。

// 序列化：将结构体序列化为字节切片，然后 反序列化：将字节切片反序列化为结构体。
type Foo struct {
	// 字段标签是赋给struct字段的元数据元素，用于定义编组和解组逻辑如何查找和处理关联的元素。
	Bar string `xml:"id,attr"`
	Baz string `xml:"parent>child"`
}

func main() {
	f := Foo{"Joe Junior", "Hello Senior"}
	// 序列化：将结构体序列化为字节切片
	b, _ := json.Marshal(f) // Marshal() method encodes the struct to JSON, return a byte slice
	fmt.Printf("%T\n", b)
	fmt.Println(string(b)) // a JSON-encoded string representation: {"Bar":"Joe Junior","Baz":"Hello Senior"}

	//将字节切片反序列化为结构体
	json.Unmarshal(b, &f)
	fmt.Printf("%T\n", f)
}
