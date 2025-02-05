package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "测试字符串匹配问题"
	fmt.Println(strings.Contains(text, "匹配"))   // true
	fmt.Println(strings.Contains(text, "接口"))   // false
}
