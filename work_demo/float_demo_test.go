package main

import (
	"testing"
)

//go test float_demo_test.go
// go test -v float_demo_test.go

func TestInit(t *testing.T){
	p := 9.7
	k := 100
	v := p * float64(k)
	//t.Log(reflect.TypeOf(v))
	t.Log(v)                 // error print
	t.Logf("%f", v)   // correct print
}
