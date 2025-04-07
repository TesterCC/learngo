package main

import (
	"fmt"
	"runtime"
)

/*
打印错误堆栈信息方法2:
2. runtime 获取调用栈
可以使用 runtime.Caller() 来捕获堆栈信息：
*/

func captureError() error {
	// Attention：可获取具体文件和行号
	_, file, line, _ := runtime.Caller(1) // 获取调用位置
	return fmt.Errorf("错误发生在 %s:%d", file, line)
}

func main() {
	err := captureError()
	fmt.Println(err)
}
