package main

import (
	"encoding/base64"
	"fmt"
)

// base64 code and encode demo
func main() {
	plainText := "测试base64字符串处理"
	fmt.Println("Origin string:", plainText)

	encodeStr := base64.StdEncoding.EncodeToString([]byte(plainText))
	fmt.Println("Encoded string:", encodeStr)

	str := "5rWL6K+VYmFzZTY05a2X56ym5Liy5aSE55CG"
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	fmt.Println("Decoded Bytes and err: ")
	fmt.Println(decodeBytes, err)
	fmt.Println("Decoded string: ", string(decodeBytes))
}
