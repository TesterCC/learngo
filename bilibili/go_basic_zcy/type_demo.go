package main

import (
	"fmt"
)

// go语言type关键字应用场景总结 https://www.bilibili.com/video/BV16ZT6znE1b

const (
	Black int = iota + 30
	Red
	Green
	Yell
	Blue
	Magenta
	Cyan
	White
)

func Print(s string, color int) {
	// basic function

	//// 一个单参数的格式，%d 会被替换为一个颜色代码; 这种格式只设置文本颜色，不改变文本的其他属性
	//fmt.Printf("\x1b[%dm%s\x1b[0m\n", color, s)

	// 一个双参数的格式，1 表示加粗 / 高亮，%d 仍然表示颜色代码；这种格式会同时设置文本为加粗 / 高亮，并改变文本颜色
	fmt.Printf("\x1b[1;%dm%s\x1b[0m\n", color, s)

}

// example 01  借助 type 关键字，能够基于已有的类型创建新类型，这样可以增强代码的可读性。
type Color int

// example 03 能力升级，实现类型方法
func (color Color) Print(s string) {
	fmt.Printf("\x1b[1;%dm%s\x1b[0m\n", color, s)
}

func main() {
	Print("Test Hello World", Green)
	//Print("Test Hello World", Cyan)
	//Print("Test Hello World", Red)

	color := Color(Cyan)
	color.Print("Test2 Hello World")

	//// example 02 与众不同
	////ctx := context.WithValue(context.Background(), 35, 35)  // warning but 1.22 can't see
	//ctx := context.WithValue(context.Background(), color, color)
	//<-ctx.Done()

	child := Child{"Chaoyang"}
	fmt.Println(child.Address())

}

// example 04 化长为短，化繁为简
// f2 参数是一种函数类型，接受3个输入，返回2个输出
//func f1(a int, f2 func(bool, string, int) (int, error)) (int, error) {
//	return f2(true, "hello", a)
//}
//
//func f3(a bool, f2 func(bool, string, int) (int, error)) (int, error) {
//	return f2(true, "hello", 3)
//}

type FK func(bool, string, int) (int, error)

func f1(a int, f2 FK) (int, error) {
	return f2(true, "hello", a)
}

func f3(a bool, f2 FK) (int, error) {
	return f2(true, "hello", 3)
}

// example 05 继承
// Parent 有一个成员变量和成员方法
type Parent struct {
	Province string
}

func (Parent) City() string {
	return "Beijing"
}

type Child Parent // 继承 Parent

func (c Child) Address() string {
	// Child不能直接访问父类的成员方法，需要强制类型转化为父类才能使用。
	// 同时Child还可以有自己额外的成员方法
	return c.Province + ", " + Parent(c).City()
}
