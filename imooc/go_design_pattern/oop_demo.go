package main

import "fmt"

type Duck struct {
	color string // 首字母小写表示私有属性
	Name  string // 首字母大写表示公有属性
}

// 函数首字母大写表示方法对外公开
func (d *Duck) Move() string {
	return d.color + "的鸭子在走"
}

// 封装
type Person struct {
	name string
}

func (p *Person) Run() {
	fmt.Println(p.name + "在跑步")
}

//继承
type Chinese struct {
	p    Person
	skin string
}

func (c Chinese) GetSkin() string {
	return "Yellow Skin"
}

//多态
type Human interface {
	Speak()
}

type American struct {
	name     string
	Language string
}

type Janpanese struct {
	name     string
	Language string
}

func (a *American) Speak() {
	fmt.Println("美国人" + a.name + "说" + a.Language)
}
func (j *Janpanese) Speak() {
	fmt.Println("日本人" + j.name + "说" + j.Language)
}

func main() {
	d := Duck{"Yellow", "David"}
	fmt.Println(d.Move())

	p := Person{"Patrick"}
	p.Run()

	//ch := Chinese{Person{"Pang"}, ""}
	ch := Chinese{p, ""}
	fmt.Println(ch.p.name + "'s skin is " + ch.GetSkin())

	a := American{"Jack","English"}
	a.Speak()

	j := Janpanese{"Mori", "Japanese"}
	j.Speak()
}
