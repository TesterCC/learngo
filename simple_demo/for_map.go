package main

import "fmt"

func main() {
	a := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}

	// go开发者认为 map 应该有随机性, 所以故意加了随机初始化 位置的函数 确保 map 每次遍历都是从不同的位置开始
	// 遍历的顺序是 相同的, 只是初始位置不同
	doRange(a)
	doRange(a)
	doRange(a)
	doRange(a)

}

func doRange(a map[int]int){
	for k,_ := range a{
		fmt.Print(k,"|")
	}
	fmt.Println()
}