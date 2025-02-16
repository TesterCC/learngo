package main

import (
	"fmt"
	"time"
)

/*
https://tour.go-zh.org/flowcontrol/12
https://go.dev/tour/flowcontrol/12

defer 推迟

defer 语句会将函数推迟到外层函数返回之后执行。

推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。

go 中提供了一个 defer 语句用来延迟一个函数(匿名函数)或者方法的执行，它会在函数执行完成(return)之后调用。
defer 语句的执行顺序遵循 后进先出（LIFO） 的原则，如果有多个 defer 语句，即最后定义的 defer 语句会最先执行。

一般为了防止代码里有资源泄露， 对于打开的资源比如文件等我们需要显式关闭，这种场合就是 defer 发挥作用最好的场景，也是 go代码中使用 defer 最常用的场景。

*/

func main() {

	defer fmt.Println("world!!!")                        // 5
	defer fmt.Printf("Default format: %v\n", time.Now()) // 4  because defer is LIFO

	fmt.Println("Hello") // 1

	fmt.Printf("Default format: %v\n", time.Now()) // 2

	fmt.Println("Normal output end ...") // 3

}
