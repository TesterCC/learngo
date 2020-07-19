package main

import "fmt"

/*
todo current site
http://127.0.0.1:3999/methods/16
https://tour.go-zh.org/methods/16

类型选择 是一种按顺序从几个类型断言中选择分支的结构。

类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）， 它们针对给定接口值所存储的值的类型进行比较。

类型选择中的声明与类型断言 i.(T) 的语法相同，只是具体类型 T 被替换成了关键字 type。

此选择语句判断接口值 i 保存的值类型是 T 还是 S。在 T 或 S 的情况下，变量 v 会分别按 T 或 S 类型保存 i 拥有的值。在默认（即没有匹配）的情况下，变量 v 与 i 的接口类型和值相同。
*/

func do(i interface{})  {
	switch v := i.(type){
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case bool:
		fmt.Printf("%v is Bool Type\n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

}

func main() {

	do(7)
	do("Test")
	do(true)
	do(13.233333333)
}
