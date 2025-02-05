package main

import (
	"fmt"
	"reflect"
)

// P81 一般方法的反射调用非常简单，直接使用 reflect.ValueOf 根据函数指针获取方法的Value即可。

func main() {
	methodOfHello := reflect.ValueOf(hello)
	methodOfHello.Call([]reflect.Value{})
}

func hello() {
	fmt.Print("Hello Go World!")
}
