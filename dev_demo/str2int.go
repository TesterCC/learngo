package main

import (
	"fmt"
	"strconv"
)

// ref: https://www.cnblogs.com/liujie-php/p/10558641.html

// go语言string、int、int64互相转换
// http://golang.org/pkg/fmt/

var i int = 100
var i2 int64 = 222.00
var x string = "7777777"


func main() {
	fmt.Println("1. int转str: ")
	// int 转 str,   %d代表Integer
	// 通过Itoa方法转换
	str1 := strconv.Itoa(i)

	// 通过Sprintf方法转换
	str2 := fmt.Sprintf("%d", i)

	// int64到string
	fmt.Printf("i2 type is: %T\n", i2)
	str3 := strconv.FormatInt(i2,10)

	fmt.Println("int通过Itoa方法转str: ", str1)
	fmt.Println("int通过Sprintf方法转换str: ", str2)
	fmt.Println("int64通过FormatInt方法转换str: ", str3)
	fmt.Printf("%T,%T,%T\n", str1, str2, str3)

	fmt.Println("==================================")

	fmt.Println("2. string转int: ")
	// string to int
	// string到int
	int,_:=strconv.Atoi(x)
	fmt.Printf("str转换后的int type: %T\n", int)
	fmt.Println("str通过Atoi转换为int: ", int)

	// string到int64
	int64, _ := strconv.ParseInt(x, 10, 64)
	fmt.Printf("str转换后的int type: %T\n", int64)
	fmt.Println("str通过Atoi转换为int: ", int64)

}
