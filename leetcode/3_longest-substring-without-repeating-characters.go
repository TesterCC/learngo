package main

/*
例题：
[无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/?utm_source=LCUS&utm_medium=ip_redirect_q_uns&utm_campaign=trans)
[英文版](https://leetcode.com/problems/longest-substring-without-repeating-characters)
同：[剑指Offer 48 寻找最长不含有重复字符的子串](https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/)

解题思路：

对于每一个字母x
- `lastOccurred[x]`不存在，或者`< start` --> 无需操作
- `lastOccurred[x] >= start ` --> 更新start
- 更新`lastOccurred[x]`，更新maxLength
*/

func lengthOfNonRepeatingSubStr(s string) int {
	// 哈希集合，记录每个字符是否出现过
	// todo

}

func main() {
	
}
