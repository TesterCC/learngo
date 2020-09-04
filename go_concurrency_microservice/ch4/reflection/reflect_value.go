package main

import (
	"fmt"
	"reflect"
)

// P78-79 反射值对象
func main() {
	name := "Alice"

	valueOfName := reflect.ValueOf(name)

	fmt.Println(valueOfName.Interface())
	// 因此再不清楚变量的具体类型时，建议先使用Value.Interface方法取值，再通过类型推导进行赋值。

	// 如果取值变量的类型与取值的方式不匹配，那么程序就会panic
	//fmt.Println(valueOfName.Bytes())  // uncomment will cause panic

	// 对变量的修改可以通过Value.Set方法实现
	valueOfName2 := reflect.ValueOf(&name)  // 获取name指针的value
	valueOfName2.Elem().Set(reflect.ValueOf("Bob"))
	fmt.Println(name)

	name3 := "Chris"
	valueOfName3 := reflect.ValueOf(name3)
	fmt.Printf("name can be address : %t\n", valueOfName3.CanAddr())
	valueOfName4 := reflect.ValueOf(&name3)
	fmt.Printf("name can be address : %t\n", valueOfName4.CanAddr())
	valueOfName5 := valueOfName4.Elem()
	fmt.Printf("name can be address : %t\n", valueOfName5.CanAddr())   // 只有指针类型解引后的Value才是可寻址的
}
