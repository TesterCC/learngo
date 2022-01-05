package main

import "fmt"

/*
19 - Golang中面向对象继承
https://www.bilibili.com/video/BV1gf4y1r79E?p=19
*/

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat() ...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk() ...")
}

// ======================================
// 1.继承
type SuperMan struct {
	Human  // 重点：SuperMan类继承了Human类的方法

	level int
}

// 2.重定义父类的方法Eat()
func (this *SuperMan) Eat()  {
	fmt.Println("SuperMan.Eat() ...")
}

// 3.子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly() ...")
}


func (this *SuperMan) Print() {
	fmt.Println("Name = ", this.name)
	fmt.Println("Ad = ", this.sex)
	fmt.Println("Level = ", this.level)
}

func main() {
	h := Human{"Alice", "female"}

	h.Eat()
	h.Walk()

	fmt.Println("--------------------------------------")
	// 定义一个子类对象
	//s := SuperMan{Human{"Cloud", "male"}, 38}

	// 另一种方式定义子类对象
	var s SuperMan
	s.name = "Alan"
	s.sex = "male"
	s.level = 99

    s.Walk()  // 父类的方法
    s.Eat()   // 子类的方法
    s.Fly()   // 子类的方法

    s.Print()
}
