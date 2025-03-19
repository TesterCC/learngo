package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 使用 os.Mkdir() 创建单个目录, 如果目录的父目录不存在，会返回错误。如果目录已存在，也会报错。
	err := os.Mkdir("./mydir", 0755)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created directory: mydir")

	// 推荐：使用 os.MkdirAll() 创建多级目录，递归创建目录及其所有父目录，如果目录已经存在，不会返回错误。
	err = os.MkdirAll("./parent/child/grandchild", 0755)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created directory: parent/child/grandchild")

	// 使用 os.MkdirTemp() 创建临时目录，目录名称是随机生成的，确保唯一性。
	dir, err := os.MkdirTemp("./", "mytempdir-")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created temporary directory:", dir) // e.g. mytempdir-3811348933

	// 清理临时目录
	defer os.RemoveAll(dir)
}
