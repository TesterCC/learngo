package main

import "fmt"

/*
https://pegasuswang.github.io/LetsGo/basics/06_error/error/
根据自己的业务代码定义error

只需要自己定义一个结构体， 然后实现 Error() 方法就实现了 go 的 error 接口

e.g. 服务为Article服务，自定义ArticleError的错误类型
可以直接参考 errors.go
*/

type ArticleError struct {
	Code int32
	Message string
}

func (e *ArticleError) Error() string {
	return fmt.Sprintf("[ArticleError] Code=%d, Message=%s", e.Code, e.Message)
}

func NewArticleError(code int32, message string) error {
	return &ArticleError{
		Code: code,
		Message: message,
	}
}

// 如果出现了业务错误，你就可以调用 NewArticleError 函数并且传入你业务里定义的错误码和错误信息创建一个错误类型了。
func TriggerArticleError(s string) (int, error){
	if len(s) == 0 {
		return 0, NewArticleError(10001, "Article string is empty.")
	}
	return 1,nil
}

func main() {
	res, err := TriggerArticleError("")   // trigger error
	fmt.Println("1.", res)
    if err != nil {
    	fmt.Println(err)
	}

}
