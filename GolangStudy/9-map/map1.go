package main

import "fmt"

func main() {
	// 第1种声明方式
	// 声明myMap1是一种map类型 key是string， value是string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 is empty map")
	}

	// 在使用map前，需要先用make给map分配数据空间；不分配数据空间就赋值会引发panic
	myMap1 = make(map[string]string,10)
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"
	myMap1["four"] = "go"
	myMap1["five"] = "rust"

    fmt.Println("myMap1: ", myMap1)

	// 第2种声明方式  可以不用指定开辟多少空间，默认可开辟  一般开发中无初始化用这种
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"

	fmt.Println("myMap2: ", myMap2)

	// 第3种声明方式 直接创建map并初始化   一般开发中有初始化用这种
	myMap3 := map[string]string {
		"one": "php",
		"two": "c++",
		"three": "python",
	}

	fmt.Println("myMap3: ", myMap3)


}
