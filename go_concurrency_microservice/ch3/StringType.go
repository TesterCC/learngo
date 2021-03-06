package main

import (
	"fmt"
	"unicode/utf8"
)

// P44-45
/*
字符出型，基于UTF-8编码实现，需要区分byte和rune

byte，类型为uint8，代表了一个ASCII字符。
rune，类型为int32，代表一个UTF-8字符。  中文字符在UTF-8中占用了3个字节。 int32能表达更多的值，更容易处理Unicode字符。
*/

func main() {
	f := "Golang编程"
	fmt.Printf("byte len of is %v\n", len(f))                    // 12
	fmt.Printf("rune len of is %v\n", utf8.RuneCountInString(f)) // 8  统计字符串的Unicode字符数量

	// 按字节byte遍历字符串
	for _, g := range []byte(f) {
		fmt.Printf("%c", g)  // 输出为： Golangç¼ ç¨ %    原因：按byte遍历时，中文字符的Unicode字符会被截断，导致中文字符输出乱码。
	}
}
