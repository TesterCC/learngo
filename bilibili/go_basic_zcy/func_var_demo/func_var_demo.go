package main

import "fmt"

/*
对比go和python函数传参的底层逻辑
https://www.bilibili.com/video/BV1u33TzEE83/
*/

func foo(n int) {
	fmt.Printf("n=%d, n的地址：%p\n", n, &n)
	n++
	fmt.Printf("n=%d, n的地址：%p\n", n, &n)
}

func main() {
	i := 8
	fmt.Printf("i=%d, i的地址：%p\n", i, &i)
	foo(i) // go语言这里实际上是把备份传给了foo函数，所以函数内的变量和函数外的变量地址不同。
	fmt.Printf("i=%d, i的地址：%p\n", i, &i)
}

/*
i=8, i的地址：0xc00000a0f8
n=8, n的地址：0xc00000a118
n=9, n的地址：0xc00000a118
i=8, i的地址：0xc00000a0f8
*/
