package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

/*

http://127.0.0.1:3999/moretypes/23
https://tour.go-zh.org/moretypes/23


练习：映射
实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。

你会发现 strings.Fields 很有帮助。
https://go-zh.org/pkg/strings/#Fields

strings.Fields()
以连续的空白字符为分隔符，将 s 切分成多个子串，结果中不包含空白字符本身
Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning an array of substrings of s or an empty list if s contains only white space.
*/

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	//fmt.Println(words)
	ret := make(map[string]int)
	for _,w := range(words){
		ret[w] += 1
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}


