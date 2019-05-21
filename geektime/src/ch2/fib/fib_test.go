package fib

import (
	"testing"
)

/*
实现Fibonacci数列：
1,1,2,3,5,8,13,...
*/
// Go更推荐统一的赋值方式，Go 2.0可能会取消一些赋值方式

func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int = 1

	//var (
	//	a = 1
	//    b = 1
	//    )

	// simplest
	a := 1
	b := 1

	//fmt.Print(a)
	t.Log(a)

	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}

}

func TestExchange(t *testing.T) {
	a := 1
	b := 7
	//tmp:=a
	//a=b
	//b=tmp
	a, b = b, a
	t.Log(a, b)
}
