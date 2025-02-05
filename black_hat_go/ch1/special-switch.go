package main

import "fmt"

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

// 这个更直观
func showAnimal(animal AnimalIF) {
	animal.Sleep() // 多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}



// P15 special syntax
func foo(i interface{}) {
	// retrieve
	switch v := i.(type) {
	case int:
		fmt.Println("I'm an integer!")
	case string:
		fmt.Println("I'm a string!")
	default:
		fmt.Println("Unknown type!")
		//_ = v
		fmt.Println(v)
	}

}

func main() {
	// need to learn go interface to understand

	cat := Cat{"Black"}
	showAnimal(&cat)
	foo(cat)
}
