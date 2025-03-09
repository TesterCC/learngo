package main

import (
	"fmt"
	"strconv"
)

// python3 demo: str2hex.py

func main() {
	s := "/bin/sh"
	var hexVals []string
	for _, char := range s {
		asciiVal := int(char)

		// 无符号 8 位整数（byte）类型的切片，其中每个元素是实际的十六进制数值。
		hexVal := strconv.FormatInt(int64(asciiVal), 16)
		hexVals = append(hexVals, "0x"+hexVal)
	}
	fmt.Println(hexVals)
	// [0x2f 0x62 0x69 0x6e 0x2f 0x73 0x68]
}
