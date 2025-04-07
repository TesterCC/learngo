package main

import "fmt"

func main() {
	// 打印大写字母
	fmt.Println("大写字母:")
	// 本质是基于字符的 ASCII 码值开展数值遍历
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Println()

	// 打印小写字母
	fmt.Println("小写字母:")
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Println()
}
