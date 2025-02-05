package main

import "fmt"

//  切片的截取
func main() {
	s := []int{1,2,3}  // len = 3, cap = 3, [1,2,3]

	// 左闭右开  [0,2)
	s1 := s[0:2]      // [1,2]
	fmt.Println("s1 = ", s1)

	// s1 和 s 指向的底层数组是一致的
	s1[0] = 100
	fmt.Println(s)
	fmt.Println(s1)

	// copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int,3)  // s2 = [0,0,0]

	// 将s中的值依次copy到s2中
	copy(s2,s)    // 深拷贝
	fmt.Println("s2 = ", s2)
}
