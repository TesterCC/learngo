package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

//P19 1.2.5 可以将字符串转换为结构体，也可以将结构体转换为字符串。

// 序列化：将结构体序列化为字节切片，然后 反序列化：将字节切片反序列化为结构体。
// 自定义结构体 Foo
type Foo struct {
	// 字段标签是赋给struct字段的元数据元素，用于定义编组和解组逻辑如何查找和处理关联的元素。
	Bar string `xml:"id,attr"`      // 解析 XML 属性 id
	Baz string `xml:"parent>child"` // 解析 <parent><child> 的值
}

func main() {
	// 序列化：将结构体序列化为字节切片

	f := Foo{
		Bar: "Joe Junior",
		Baz: "Hello TesterCC",
	}

	output, err := xml.MarshalIndent(f, "", "  ") // 美化输出
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}

	fmt.Printf("%T\n", output) // debug，[]uint8
	fmt.Println(string(output))

	fmt.Println(strings.Repeat("-", 33))
	// 反序列化

	// out value same
	//data := `<Foo id="Joe Junior"><parent><child>Hello TesterCC</child></parent></Foo>`

	err2 := xml.Unmarshal(output, &f)
	if err2 != nil {
		fmt.Println("解析 XML 失败:", err2)
		return
	}

	//fmt.Printf("%T\n", f)      // debug， main.Foo
	fmt.Println("Bar:", f.Bar) // 输出: Bar: Joe Junior
	fmt.Println("Baz:", f.Baz) // 输出: Baz: Hello TesterCC
}
