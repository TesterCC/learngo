package main

import "fmt"

/*
18 - Golang中面向对象类的表示与封装
ref: https://www.bilibili.com/video/BV1gf4y1r79E?p=18

类名、属性名、⽅法名 ⾸字⺟⼤写表示对外(其他包)可以访问，否则只能够在本包内访问
*/

// 如果类名首字母大写，表示其它包也能够访问
type Hero struct {
	// 如果说类的属性首字母大写，表示该属性是对外能够访问的，否则的话只能够类的内部访问
	Name string
	Ad   int
	Level int
}

/* // 这段改进为下面的
func (hero Hero) Show() {
	fmt.Println("Name = ", hero.Name)
	fmt.Println("Ad = ", hero.Ad)
	fmt.Println("Level = ", hero.Level)
}

func (hero Hero) GetName() string {

	return hero.Name
}

func (hero Hero) SetName(newName string) {
	// hero 是调用该方法的对象的一个副本（拷贝）
	hero.Name = newName
}
*/

func (hero *Hero) Show() {
	fmt.Println("Name = ", hero.Name)
	fmt.Println("Ad = ", hero.Ad)
	fmt.Println("Level = ", hero.Level)
}

func (hero *Hero) GetName() string {

	return hero.Name
}

func (hero *Hero) SetName(newName string) {
	// hero 是调用该方法的对象的地址
	hero.Name = newName
}


func main() {
	// create a object
	hero := Hero{Name: "Alice", Ad: 100, Level: 1}

    hero.Show()

	hero.SetName("Bob")

	hero.Show()
}
