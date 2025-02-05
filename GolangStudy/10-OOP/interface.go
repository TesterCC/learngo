package main

import "fmt"

/*
20-Golang中面向对象多态的实现与基本要素
ref: https://www.bilibili.com/video/BV1gf4y1r79E?p=20

多态：父指针指向子类对象(父类引用指向子类对象)

多态的基本要素：
1.有一个父类（有接口）
2.有子类（实现了父类的全部接口方法）
3.父类类型的变量（或指针） 指向（引用） 子类的具体数据变量
*/

// 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string // get animal color
	GetType() string  // get animal type

}

// 具体的类   子类实现了接口的所有方法，就等于实现了接口
type Cat struct {
	color string // cat color
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string // cat color
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

// 这个更直观
func showAnimal(animal AnimalIF) {
	animal.Sleep() // 多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}

func main() {
	/*
		var animal AnimalIF // 接口的数据类型，即父类指针
		animal = &Cat{"Green"}
		animal.Sleep()  // 调用 Cat的Sleep()方法

		animal = &Dog{"Yellow"}
		animal.Sleep()  // 调用Dog的Sleep()方法，多态的现象
	*/

	cat := Cat{"Black"}
	dog := Dog{"Brown"}

	showAnimal(&cat)
	fmt.Println("================================")
	showAnimal(&dog)
}
