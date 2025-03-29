package main

import (
	"fmt"
	"io"
	"os"
)

func openFile(filename string) (*os.File, error) {
	return os.Open(filename) // 成功返回 *os.File，失败返回 error
}

func main() {
	// 调用 openFile 函数
	file, err := openFile("dev_demo/test.txt")

	// 错误处理
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}

	// 记得关闭文件
	defer file.Close()

	fmt.Println("文件打开成功:", file.Name())

	// 读取文件内容
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	// 打印文件内容
	fmt.Println("文件内容:")
	fmt.Println(string(content))

}
