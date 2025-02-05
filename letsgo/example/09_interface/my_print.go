package main

import "fmt"

/*

使用空接口实现泛型

ref:
https://pegasuswang.github.io/LetsGo/basics/09_interface/interface/#_2

go 目前没有直接提供对泛型的支持，学了接口之后其实我们可以用接口来实现。

如果一个 struct 实现了一个接口声明所有方法，我们就说这个 struct (隐式)实现了这个接口, 所以：
所有类型都实现了空接口(interface{})，所以可以用空接口+类型断言转成任何我们需要的类型。


如何实现泛型呢？空接口其实给了我们思路。既然它能转成所有类型，那我们以空接口作为参数不就好了嘛，这个想法是对的。
如果你有留意的话，到现在我们的代码示例里边使用最多的是啥，其实是这句话 fmt.Println()，不知道你之前有没有发现这个函数 居然可以传递任意类型进去
*/

// 实现一个简单的可以打印多种类型的 MyPrint 函数
func MyPrint(i interface{}) {
	switch o := i.(type) {
	case int:
		fmt.Printf("%d\n", o)
	case float64:
		fmt.Printf("%f\n", o)
	case string:
		fmt.Printf("%s\n", o)
	default:
		fmt.Printf("%+v\n", o)
	}
}

func main() {
	MyPrint(1)
	MyPrint(4.2)
	MyPrint("hello")
	MyPrint(map[string]string{"hello": "go"})
}

// 空接口在实现泛型的时候很有用，不过一般情况下如果不是必要，我们还是单独实现对应类型的函数就好，代码可读性也更高。


