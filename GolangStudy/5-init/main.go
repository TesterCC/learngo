package main

import (
	"learngo/GolangStudy/5-init/lib1"
	"learngo/GolangStudy/5-init/lib2"
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
}


/*
ref: https://www.bilibili.com/video/BV1gf4y1r79E?p=8

$ go run main.go
lib1, init() ...
lib2, init() ...
lib1Test() ...
lib2Test() ...

*/