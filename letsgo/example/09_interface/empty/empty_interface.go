package main

import "fmt"

/*

使用空接口实现泛型

ref:
https://pegasuswang.github.io/LetsGo/basics/09_interface/interface/#_2

go 目前没有直接提供对泛型的支持，学了接口之后其实我们可以用接口来实现。

如果一个 struct 实现了一个接口声明所有方法，我们就说这个 struct (隐式)实现了这个接口, 所以：
所有类型都实现了空接口(interface{})，所以可以用空接口+类型断言转成任何我们需要的类型。
*/

// 创建了一个空接口数组，它的元素可以是任何类型


type Dog struct {
	Name string
}

func (d Dog) Sleep() {
	fmt.Printf("Dog %s is sleeping.\n", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Sleep() {
	fmt.Printf("Cat %s is sleeping.\n", c.Name)
}

func main() {
	// 用空接口+类型断言转成任何我们需要的具体类型
	animalList := []interface{}{Dog{Name: "David"}, Cat{Name: "Tom"}}
	for _, s := range animalList {
		if dog, ok := s.(Dog); ok {
			fmt.Printf("I am a Dog, my name is %s\n", dog.Name)
		}
		if cat, ok := s.(Cat); ok {
			fmt.Printf("I am a Cat, my name is %s\n", cat.Name)
		}
	}
}

