package main

import "fmt"

// P11
type Person struct {
	Name string
	Age  int
}

type Dog struct {
}

// P12 struct and interface
// 任何实现了方法SayHello()的类型都是Friend，接口Friend实际上并未实现这个函数
type Friend interface {
	SayHello()
}

// method
func (d *Dog) SayHello() {
	fmt.Println("Woof woof")
}

// 函数Greet()以接口Friend作为参数并以Friend特定的方式执行，可将任何Friend类型作为参数
// 名为Gree()的函数期望将Friend类型作为参数，因此可以向其传递Person。
// PS：因为结构体Person具有方法SayHello()，可以被视为Friend类型
func Greet(f Friend) {
	f.SayHello()
}

// 给结构体Person的变量p定义一个方法SayHello()
func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}

func main() {
	var guy = new(Person)
	guy.Name = "Dave"
	//guy.SayHello()
	Greet(guy)

	var dog = new(Dog)
	Greet(dog)
}
