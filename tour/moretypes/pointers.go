package main

import "fmt"

/*
https://tour.go-zh.org/moretypes/1
https://go.dev/tour/moretypes/1

指针

p.s. go语言的指针比c语言简单很多

Go 拥有指针。指针保存了值的内存地址。

类型 *T 是指向 T 类型值的指针，其零值为 nil。

e.g.:  var p *int

& 操作符会生成一个指向其操作数的指针。

e.g.:
i := 42
p = &i

* 操作符表示指针指向的底层值。
fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i

这也就是通常所说的「解引用」或「间接引用」。

与 C 不同，Go 没有指针运算。

& 取址
* 取值

*/

func main() {
	i, j := 42, 2701

	p := &i         // 指向i , & 操作符会生成一个指向其操作数的指针，即生成指向i的指针  取址
	fmt.Println(*p) // 通过指针读取 i 的值  42

	*p = 21        // 通过指针设置 i 的值 , * 操作符表示指针指向的底层值    取值
	fmt.Println(i) // 查看 i 的值，现在应该是修改后的   21

	p = &j          // 指向j， &生成一个指向j的指针
	fmt.Println(p)  // 这里打印成出的是指针p指向的j的值的地址：0x1400000e0b0
	fmt.Println(*p) // 通过指针读取j的值  2701
	*p = *p / 37    // 通过指针对 j 进行除法运算   2701/37=73
	fmt.Println(j)  // 查看 j 的值   73
}
