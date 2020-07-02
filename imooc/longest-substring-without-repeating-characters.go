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

*/

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)   // 当前值最后出现的位置
	start := 0
	maxLength := 0  // 目前substr的最大长度

	for i,ch := range []byte(s) {   // 中文字符每个算3个
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
	ret := lengthOfNonRepeatingSubStr("testingtaaaaaa")   // esting
	fmt.Println(ret)

	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("x"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))

	// 这里有个问题，这样写只能支持英文
}
