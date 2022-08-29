package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//ebook P114

// 允许学生输入他们的分数，并告诉他们是否通过。 >60% 算通过
// pass_fail reports whether a grade is passing or failing.

func main() {
	fmt.Print("Enter a grade: ")
	// 设置从键盘获取文本的"缓冲读取器"，即从程序的标准输入中读取输入内容。然后返回新的 bufio.Reader
	reader := bufio.NewReader(os.Stdin)
	// 返回用户输入的所有内容，知道按下Enter键为止。
	// 处理返回错误值，将其存储在变量中
	input, err := reader.ReadString('\n')
	// 报告错误并停止运行  Fatal 将一条消息记录到终端并停止程序运行
	if err != nil {
		log.Fatal(err)
	}
	// 打印用户输入的内容
	fmt.Println(input)

	if true {
		fmt.Println("I'll be printed!")
	}

	if false {
		fmt.Println("I won't be printed!")
	}

	input = strings.Replace(input, "\n", "", -1)
	grade, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	if grade == 100 {
		fmt.Println("Perfect!")
	} else if grade >= 60 {
		fmt.Println("You pass!")
	} else {
		fmt.Println("You fail!")
	}
}
