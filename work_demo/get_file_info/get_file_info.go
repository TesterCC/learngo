package main

import (
	"fmt"
	"path"
)

func main() {
	// 提取文件名
	filePath := "./t1.txt"
	// 用于从文件路径中提取最后一个部分（即文件名或目录名）。它的作用是返回路径的最后一个元素，类似于 Unix 中的 basename 命令。
	filename := path.Base(filePath)
	fmt.Println(filename) // t1.txt

	// example
	// 提取文件名
	fmt.Println(path.Base("/home/user/file.txt")) // 输出: file.txt

	// 提取目录名
	fmt.Println(path.Base("/home/user/tester")) // 输出: tester

	// 处理根路径
	fmt.Println(path.Base("/")) // 输出: ""

	// 处理空路径
	fmt.Println(path.Base("")) // 输出: .

	// 处理相对路径
	fmt.Println(path.Base("dir/file.txt")) // 输出: file.txt
}
