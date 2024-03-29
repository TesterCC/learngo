package main

import (
	"bufio"
	"fmt"
	"os"
)

//ebook P114

// 允许学生输入他们的分数，并告诉他们是否通过。 >60% 算通过
// pass_fail reports whether a grade is passing or failing.

func main() {
	fmt.Print("Enter a grade: ")
	// 设置从键盘获取文本的"缓冲读取器"，即从程序的标准输入中读取输入内容。然后返回新的 bufio.Reader
	reader := bufio.NewReader(os.Stdin)
	// 返回用户输入的所有内容，知道按下Enter键为止。
	// method1:使用空白标志符忽略错误返回值
	input, _ := reader.ReadString('\n')
	// 打印用户输入的内容
	fmt.Println(input)
}