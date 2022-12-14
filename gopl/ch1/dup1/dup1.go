package main

import (
	"bufio"
	"fmt"
	"os"
)

// gopl P61.3 找出重复行
// 输出标准输入中出现次数大于1的行，字符：出现次数
// gopl.io/ch1/dup1

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// 注意：忽略 input.Err()中可能的错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
