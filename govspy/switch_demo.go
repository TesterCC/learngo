package main

import (
	"fmt"
	"strconv"
)

func str2int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Not a number")
	}
	return i
}

func main() {
	var number_string string
	fmt.Scanln(&number_string)
	number := str2int(number_string)

	switch number {
	case 8:
		fmt.Println("Oxygen")
	case 1:
		fmt.Println("Hydrogen")
	case 2:
		fmt.Println("Helium")
	case 11:
		fmt.Println("Sodium")
	default:
		fmt.Printf("I have no idea what %d is\n", number)
	}

	// Alternative solution

	fmt.Scanln(&number_string)
	db := map[int]string{
		1:  "Hydrogen",
		2:  "Helium",
		8:  "Oxygen",
		11: "Sodium",
	}
	number = str2int(number_string)

	// 短变量声明和条件判断:允许在条件判断时声明一个或多个变量，并且在条件语句的作用域内使用这些变量。
	// 在 Go 语言中，当你尝试从一个映射（map）中获取一个值时，如果这个值存在，它会成功返回，并且 exists 会被设置为 true；如果这个值不存在，exists 会被设置为 false。
	// 这个特性允许你在尝试获取一个值的同时检查是否成功获取了这个值，而不会因为值不存在而导致程序出现运行时错误。
	if name, exists := db[number]; exists {
		fmt.Println(name)
	} else {
		fmt.Printf("I have no idea what %d is\n", number)
	}
}