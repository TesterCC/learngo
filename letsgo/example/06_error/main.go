package main

import (
	"fmt"
)

/*
https://pegasuswang.github.io/LetsGo/basics/06_error/error/
https://pegasuswang.github.io/LetsGo/basics/06_error/error/#defer
https://www.bilibili.com/video/av91237117

go 中提供了一个 defer 语句用来延迟一个函数(匿名函数)或者方法的执行，它会在函数执行完成之后调用。
一般为了防止代码里有资源泄露， 对于打开的资源比如文件等我们需要显示进行关闭，这种场合就是 defer 发挥作用最好的场景，也是 go 代码中使用 defer 最常用的场景。

如果你用过 python 的话，go 中的 defer 和 python 使用 with 语句保证资源会被关闭目的一样。
另外函数里可以使用多个 defer 语句，如果有多个 defer 它们会按照后进先出(Last In First Out)的顺序执行。
*/

func testDefer() string {

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("function body")
	return "hello"
}

// 自定义业务异常
type ArticleError struct {
	Code    int32
	Message string
}

func (e *ArticleError) Error() string {
	return fmt.Sprintf("[ArticleError] Code=%d, Message=%s", e.Code, e.Message)
}

func NewArticleError(code int32, message string) error {
	return &ArticleError{
		Code:    code,
		Message: message,
	}
}

func MustDivide(a, b int) int {
	if b == 0 {
		panic("divide by zero") // 让程序停止运行
	}
	return a / b
}


/*
如果不幸传入了除数为0，但是又不想让进程退出呢？go 还提供了一个 recover 函数用来从异常中恢复，
比如使用 recover 可以把一个 panic 包装成为 error 再返回，而不是让进程退出
 */
func Divide2(a, b int) (res int, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()
	res = MustDivide(a, b)
	return // 命名返回值不用加上返回的参数
}



func main() {
	fmt.Println(testDefer())
	/*  result:
	function body
	defer 2  LIFO
	defer 1
	hello
	*/

	//MustDivide(3,0)
	//fmt.Println("end")

	//err := NewArticleError(10001, "Testcase erorr!")
	//fmt.Println(err)

	res, err := Divide2(10, 0)
	fmt.Println(res, err)

}
