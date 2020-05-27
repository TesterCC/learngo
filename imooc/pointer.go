package main

import "fmt"

// 2-7 实现swap函数 交换a和b的值
//func swap(a,b int){
//	b,a = a,b   // a,b是值传递，改了没用不能这样写
//}

func swap(a, b *int) {
	*b, *a = *a, *b // *a,*b是引用传递
}

// 这样实现更优雅
func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b) // 4,3

	c, d := 5, 6
	c, d = swap2(c, d)
	fmt.Println(c, d) // 6,5

}
