package main

import "fmt"

// const 来定义枚举类型
// 可以在const()添加一个关键字 iota，每行的iota都会累加1，第一行的iota默认值是0
const (
	BEIJING   = 10 * iota // iota = 0
	SHANGHAI              // iota = 1
	SHENZHEN              // iota = 2
	GUANGZHOU             // iota = 3
)

// 使用场景：权限幂值，定义枚举，定义code码
// iota 只能配合 const() 一起使用， iota只有在const()中进行累加效果。
const (
	a, b = iota + 1, iota + 2 // iota = 0, a = iota + 1 = 1, b = iota + 2 = 2
	c, d                      // iota = 1, c = iota + 1 = 2, d = iota + 2 = 3
	e, f                      // iota = 2, e = 3, f = 4

	g, h = iota * 2, iota * 3 // iota = 3, g = iota * 2 = 6, h = iota * 3 = 9
	i, j                      // iota = 4, i = 8, j = 12
)

func main() {
	// 常量（只读属性）
	const length int = 10
	fmt.Println("length = ", length)
	//length = 300      // 会报错，常量是不允许修改的

	fmt.Println("BJ = ", BEIJING)
	fmt.Println("SH = ", SHANGHAI)
	fmt.Println("SZ = ", SHENZHEN)
	fmt.Println("GZ = ", GUANGZHOU)
	fmt.Printf("BJ type is %T\n", BEIJING) // iota = 0
	fmt.Printf("SZ type is %T\n", SHENZHEN)

	fmt.Println(a, b, c, d, e, f, g, h, i, j)
}
