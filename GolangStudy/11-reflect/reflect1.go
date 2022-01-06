package main

import (
	"fmt"
	"reflect"
)

/*
23-golang反射reflect机制用法
ref: bilibili.com/video/BV1gf4y1r79E?p=23

reflect包
1. ValueOf用来获取输入参数接口中的数据的值，如果接口为空则返回0
2. TypeOf用来动态获取输入参数接口中的值的类型，如果接口为空则返回nil
*/

// 最基本的反射用法
func reflectNum(arg interface{})  {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.2345
	reflectNum(num)
}
