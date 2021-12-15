package main

import (
	"testing"
)

// go test float_demo_test.go
// go test -v float_demo_test.go

/*
Go的自动化测试框架比JUnit、PyUnit等更加轻量级。
要点如下：
- 测试代码以xxx_test.go方式命名
- 测试代码中import “testing”
- 测试函数形如 func Testxyz(t *Testing.T) {…}
- 执行测试：go test
 */

func TestInit(t *testing.T){
	p := 9.7
	k := 100
	v := p * float64(k)
	//t.Log(reflect.TypeOf(v))
	t.Log(v)                 // error print
	t.Logf("%f", v)   // correct print
}
