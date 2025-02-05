package main

import "fmt"

/*
P418 附录A: Go语言常见坑

这些常见坑都是符合go语言语法的，可以正常编译，但是可能是运行结果错误，或者是有资源泄漏的风险。
*/


func main() {
	// 1.可变参数是空接口类型，要注意参数展开问题
	a := []interface{}{1, 2, 3}

	fmt.Println(a)    // [1 2 3]
	fmt.Println(a...) // 1 2 3

	// 2.数组是值传递
	// 在函数调用参数中，数组是值传递，无法通过修改数组类型的参数返回结果, 必要时需要使用切片
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr) // [7 2 3]
	}(x)

	fmt.Println(x) // [1 2 3]

	// 3.map遍历是顺序不固定的。map是一种hash表实现，每次遍历的顺序都可能不一样。

	m := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}

	for k, v := range m {
		println(k, v)
	}


	// 5.recover必须在defer函数中运行。recover捕获的是祖父级调用时的异常，直接调用时无效
	//recover()
	//panic(1)
	defer func(){
		recover()
	}()
	panic(1)

	// 6.后台的Goroutine无法保证完成任务
	// 7.通过Sleep来回避并发中的问题也是坑：1.休眠并不能保证输出完整的字符串；2.类似的还有通过插入调度语句
	// 8.独占CPU导致其它Goroutine饿死
}


// 4.返回值被屏蔽
// 在局部作用域中，命令的返回值内同名的局部变量屏蔽
//func Foo() (err error) {
//	if err := Bar(); err != nil {
//		return
//	}
//	return
//}


