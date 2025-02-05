package main

import "fmt"

func main() {
	// create map
	m := map[string]string{
		"name":    "testercc",
		"course":  "golang",
		"site":    "github",
		"quality": "good",
	}

	fmt.Println(m)

	m2 := make(map[string]int) // m2 == empty map
	fmt.Println(m2)

	var m3 map[string]int // m3 == nil
	fmt.Println(m3)

	// map的遍历
	fmt.Println("Traversing Map")
	for k, v := range m {
		fmt.Println(k, v) // key在map中是无序的，hash map，故key的顺序每次执行可能都不同
	}

	fmt.Println("只遍历key:")
	// 只遍历key
	for k := range m {
		fmt.Println(k) // key的顺序每次执行可能都不同
	}

	fmt.Println("只遍历value:")
	// 只遍历value
	for _, v := range m {
		fmt.Println(v) // key的顺序每次执行可能都不同，一般用得不多
	}

	fmt.Println("Getting Values")
	//courseName := m["course"]
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	causeName, ok := m["cause"] // 测试写错变量名的情况
	fmt.Println(causeName, ok)  // 打印空行，zero value

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("Key does not exist...")
	}

	fmt.Println("Deleting value:")
	name, ok := m["name"]
	fmt.Println(m)
	fmt.Println(name,ok)

	delete(m, "name")
	fmt.Println(m)
}
