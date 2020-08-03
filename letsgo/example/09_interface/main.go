package main

import "fmt"

/*
https://pegasuswang.github.io/LetsGo/basics/09_interface/interface/
*/

// 隐式实现，实现了Sleep()等于实现了接口
// Sleeper 接口声明
type Sleeper interface {
	Sleep() // 声明一个 Sleep() 方法
}

type Eater interface {
	Eat(foodName string)
}

// go 的 struct 可以通过嵌入实现代码复用，go 的接口也支持嵌入
// 通过嵌入两个接口实现新的接口
type LazyAnimal interface {
	Sleeper
	Eater
}

// 定义2个struct    它们都实现了 Sleep 方法，也就是说隐式实现了 Sleeper 接口。
type Dog struct {
	Name string
}

// 给Dog添加方法
func (d Dog) Sleep() {
	fmt.Printf("Dog %s is sleeping.\n", d.Name)
}

func (d Dog) Eat(foodName string) {
	fmt.Printf("Dog %s is eating %s.\n", d.Name, foodName)
}

type Cat struct {
	Name string
}

func (c Cat) Sleep() {
	fmt.Printf("Cat %s is sleeping.\n", c.Name)
}

func (c Cat) Eat(foodName string) {
	fmt.Printf("Cat %s is eating %s.\n", c.Name, foodName)
}

// 然后我们编写一个函数，不过为了支持多态，函数的参数是一个接口类型而不是具体的 struct 类型。
func AnimalSleep(s Sleeper) {
	s.Sleep()
}

// 这个是之前的main函数
func test() {
	var s Sleeper
	dog := Dog{Name: "David"}
	cat := Cat{Name: "Tom"}

	s = dog
	AnimalSleep(s) // 使用 dog 的 Sleep()
	s = cat
	AnimalSleep(s) // 使用 cat 的 Sleep()

	// 创建一个 Sleeper 切片
	sleepList := []Sleeper{Dog{Name: "Albert"}, Cat{Name: "Kitty"}}
	for _, s := range sleepList {
		// 将 index 匿名
		s.Sleep()
	}

}

// 现在的main函数
func main() {
	sleepList := []LazyAnimal{Dog{Name: "Bob"}, Cat{Name: "Tracy"}}
	foodName := "animal food"
	for _, s := range sleepList {
		s.Sleep()
		s.Eat(foodName)

		// 如何获取传入的到底是哪种 struct 类型呢？ go 给我们提供了一种方式叫做类型断言来获取具体的类型
		// instance, ok := interfaceVal.(RealType) // 如果 ok 为 true 的话，接口值就转成了我们需要的类型
		// 类型断言 type assert
		if dog, ok := s.(Dog); ok{
			fmt.Printf("I am a Dog, my name is %s.\n", dog.Name)
		}

		if cat,ok := s.(Cat); ok {
			fmt.Printf("I am a Cat, my name is %s.\n", cat.Name)
		}
	}

}
