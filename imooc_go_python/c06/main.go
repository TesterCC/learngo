package main

import "fmt"

// 5-1
func main() {
	// 字符串基本操作

	// 1.求解字符串的长度
	// 返回的是字节长度
	// 涉及到中文问题就产生了变化
	// unicode字符集 ，存储的时候需要编码
	// utf8编码是一个动态的编码规则，能用一个字节表示英文，go语言采用的是utf8编码规则
	//var name string = "SecCodeCat:"   // 正常
	var name string = "SecCodeCat:渗透测试" // python正常
	fmt.Println(len(name))

	// 类型转换
	name_arr := []rune(name) // rune 即 int32，将其转换成数组array
	//name_arr := []int32(name)     // return same result
	fmt.Println(len(name_arr)) // 15

}
