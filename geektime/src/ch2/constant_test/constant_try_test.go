package constant_test

import "testing"

// go test -v constant_try_test.go

const (
	Monday = iota + 1 //连续常量推荐这种写法
	Tuesday
	Wednesday
	Thursday
)

const (
	Readable = 1 << iota
	Writable
	Exxcutable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Thursday)
}

func TestConstantTry1(t *testing.T) {
	//a := 7   //0111
	a := 1 //0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Exxcutable == Exxcutable)
}
