package main

import "fmt"

/*
https://tour.go-zh.org/moretypes/22
https://go.dev/tour/moretypes/22

修改映射
在映射 m 中插入或修改元素：

m[key] = elem

获取元素：

elem = m[key]

删除元素：

delete(m, key)
通过双赋值检测某个键是否存在：

elem, ok = m[key]

若 key 在 m 中，ok 为 true ；否则，ok 为 false。

若 key 不在映射中，那么 elem 是该映射元素类型的零值。

同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。

注 ：若 elem 或 ok 还未声明，你可以使用短变量声明：

elem, ok := m[key]

*/

func main() {
	m := make(map[string]int)

	// 获取元素
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"]) // The value: 42

	// 修改元素
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"]) // The value: 48

	// 删除元素
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) // The value: 0

	// 通过双赋值检测某个键是否存在
	v, ok := m["Answer"]
	fmt.Println("值:", v, "是否存在？", ok)
	// 若 key 在 m 中，ok 为 true ；否则，ok 为 false。
	// 若 key 不在映射中，那么 v 是该映射元素类型的零值。
	// 同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。

	// custome
	if ok {
		fmt.Println("The map key already exist...")
	} else {
		fmt.Println("The map key don't exist")
	}

}
