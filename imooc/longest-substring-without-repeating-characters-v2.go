package main

import "fmt"

/*
https://www.bilibili.com/video/BV18Q4y1M7NV?p=13

leetcode:
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/?utm_source=LCUS&utm_medium=ip_redirect_q_uns&utm_campaign=trans
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的 长度。

思路：
对于买一个字母x
1. lastOccurred[x]不存在，或者 < start -> 无需操作
2. lastOccurred[x] >= start -> 更新start，start = lastOccurred[x] + 1
3. 最后，更新lastOccurred[x]，更新maxLength

增加对中文字符的支持
*/

func lengthOfNonRepeatingSubStrv2(s string) int {
	lastOccurred := make(map[rune]int)   // 当前值最后出现的位置
	start := 0
	maxLength := 0  // 目前substr的最大长度

	for i,ch := range []rune(s) {   // 中文字符每个算1个
		//fmt.Printf("%v,%v",i,ch)

		// 判断key是否存在
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		// 更新maxLength, substr 是从 start 开始到 i 结束
		if i-start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i  // 更新 当前值最后出现的位置
	}
	return maxLength
}

func main() {
	ret := lengthOfNonRepeatingSubStrv2("testingtaaaaaa") // esting
	fmt.Println(ret)

	fmt.Println(lengthOfNonRepeatingSubStrv2("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStrv2("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStrv2("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStrv2(""))
	fmt.Println(lengthOfNonRepeatingSubStrv2("x"))
	fmt.Println(lengthOfNonRepeatingSubStrv2("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStrv2("测试中文"))
	fmt.Println(lengthOfNonRepeatingSubStrv2("中文答题测试"))  // 中文支持
}
