package main

import "fmt"

/*
16 - map的使用方式
ref: https://www.bilibili.com/video/BV1gf4y1r79E?p=16

如何对一个map进行完全拷贝？
新定义一个map，然后依次遍历要拷贝的map给新map重新赋值即可。
*/

// map作为形参
func printMap(cityMap map[string]string)  {
	// cityMap是一个引用传递
	for key,value := range cityMap {
		fmt.Println("key = ",key, ", value = ", value)
	}
}

func ChangeValue(cityMap map[string]string)  {
	cityMap["UK"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	// add
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// traverse 遍历
    for key,value := range cityMap {
    	fmt.Println("key = ",key, ", value = ", value)
	}

	// delete
	delete(cityMap, "Japan")

    // modify
    cityMap["USA"] = "DC"
	ChangeValue(cityMap)

    fmt.Println("-----------------------------")

    printMap(cityMap)

}
