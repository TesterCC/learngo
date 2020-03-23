package main

import "fmt"

/*
http://127.0.0.1:3999/flowcontrol/12
https://tour.go-zh.org/flowcontrol/12

defer 语句会将函数推迟到外层函数返回之后执行。

推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。

go 中提供了一个 defer 语句用来延迟一个函数(匿名函数)或者方法的执行，它会在函数执行完成(return)之后调用。
一般为了防止代码里有资源泄露， 对于打开的资源比如文件等我们需要显式关闭，这种场合就是 defer 发挥作用最好的场景，也是 go 代码中使用 defer 最常用的场景。

*/

func main() {

	defer fmt.Println("world!!!")
	fmt.Println("Hello")
	
}
