package main

import (
	"fmt"
	"strconv"
)

/*
strconv.FormatInt() 是 Go 语言标准库 strconv 包中的一个函数，用于将 整数（int64 类型） 转换为 字符串。
它的功能是将整数按照指定的进制（如二进制、八进制、十进制、十六进制等）格式化为字符串。

- strconv.FormatInt() 用于将 int64 类型的整数转换为指定进制的字符串。
- 支持的进制范围：2 到 36 之间的任意进制。进制参数 base 必须在 2 到 36 之间，否则会返回空字符串。
- 适用于需要将整数格式化为字符串的场景，如日志记录、数据存储、网络传输等。
*/

func main() {
	num := int64(42)

	// 1.将整数转换为十进制字符串
	str := strconv.FormatInt(num, 10)
	fmt.Printf("str type: %T\n", str)
	fmt.Println("str value:", str) // 24

	// 2.将整数转换为二进制字符串
	str2 := strconv.FormatInt(num, 2)
	fmt.Printf("str type: %T\n", str2)
	fmt.Println("str value:", str2) // 101010

	// 3.将整数转换为八进制字符串
	str3 := strconv.FormatInt(num, 8) // 52
	fmt.Printf("str type: %T\n", str3)
	fmt.Println("str value:", str3)

	// 4.将整数转换为十六进制字符串
	str4 := strconv.FormatInt(num, 16) // 2a
	fmt.Printf("str type: %T\n", str4)
	fmt.Println("str value:", str4)

	// 255转16进制字符串
	str5 := strconv.FormatInt(255, 16) // ff
	fmt.Printf("str type: %T\n", str5)
	fmt.Println("str value:", str5)

	// 输入的整数是负数
	num2 := int64(-42)
	str6 := strconv.FormatInt(num2, 10)
	fmt.Printf("str type: %T\n", str6)
	fmt.Println("str value:", str6) // -42

}
