package main

import (
	_ "learngo/GolangStudy/5-init/lib1"      // 匿名导入，无法使用当前包的方法，但是会执行当前包内部的init()方法
	mylib2 "learngo/GolangStudy/5-init/lib2" // 别名导入，别名.方法名() 调用。实际中更推荐这个方式，不要轻易用下面的直接导入当前本包的方式
	//. "learngo/GolangStudy/5-init/lib2"    // 将当前lib2包中的全部方法，导入到当前本包的作用域中，lib2包中的全部方法可以直接使用API来调用，不需要lib2.Lib2Test()来调用
)

func main() {

	mylib2.Lib2Test()
	//Lib2Test()
}


/*
https://www.bilibili.com/video/BV1gf4y1r79E?p=9

$ go run main.go
lib1, init() ...
lib2, init() ...
lib2Test() ...

*/