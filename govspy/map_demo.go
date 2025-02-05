package main

import (
	"fmt"
	"strings"
)

func main() {
	elements := make(map[string]int)
	elements["H"] = 1
	fmt.Println(elements["H"])

	// remove by key
	elements["O"] = 8

	fmt.Println(elements)
	delete(elements, "O")
	fmt.Println(elements)

	// only do something with a element if it's in the map
	if number, ok := elements["O"]; ok {
		fmt.Println(number) // won't be printed
	}
	if number, ok := elements["H"]; ok {
		fmt.Println(number) // 1
	}

	// custom test approximate equal to python dict
	info := make(map[string]interface{})   // 最符合python dict的结构
	info["name"] = "Tester"
	info["gender"] = "Secret"
	info["addr"] = "Here"
	info["age"] = 42
	info["online"] = true

	longStr := strings.Repeat("-",33)
	fmt.Println(longStr)

	fmt.Println(len(info))
	fmt.Println(info)

	// 遍历map
	for key, value := range info {
		//fmt.Println(key, ":", value)
		fmt.Printf("key: %s, value type: %T\n", key, value)
	}
}