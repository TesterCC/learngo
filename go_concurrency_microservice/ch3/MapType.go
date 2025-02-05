package main

import "fmt"

/* P56
Map, Go语言中提供映射关系的字典，其内部通过散列表的方式实现。

make(map[keyType]valueType)
*/

func main() {
	classMates1 := make(map[int]string) // make(map[keyType]valueType)

	// 添加映射关系
	classMates1[0] = "Alice"
	classMates1[1] = "Bob"
	classMates1[2] = "Lucy"

	// 根据key获取value
	fmt.Printf("id %v is %v\n", 1, classMates1[1])

	// 在声明时初始化数据
	classMates2 := map[int]string{
		0: "Alice",
		1: "Bob",
		2: "Lucy",
	}

	fmt.Printf("id %v is %v\n", 2, classMates2[2])
	fmt.Printf("id %v is %v\n", 3, classMates2[3]) // 不存在的key，则返回valueType的默认值

	mate, ok := classMates2[1]
	//mate, ok := classMates2[4]
	if ok {
		fmt.Println(mate)
	} else {
		fmt.Println("Key not found")
	}
}
