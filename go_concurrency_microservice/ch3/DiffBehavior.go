package main

import "fmt"

// P68 3-8 内嵌不同结构体表现不同行为
// 总结：内嵌结构体特性可以实现对象的组合特性，提高代码的可复用性和可扩展性

// 游泳特性
type Swimming struct {
	
}

func (swim *Swimming) swim(){
	fmt.Println("Swimming is my ability.")
}

// 飞行特性
type Flying struct {
	
}

func (fly *Flying) fly() {
	fmt.Println("Flying is my ability.")
}

// 野鸭，具备飞行和游泳的特性
type WildDuck struct {
	Swimming
	Flying
}

// 家鸭，只具备游泳特性
type DomesticDuck struct {
	Swimming
}

func main() {
	// 声明一只野鸭
	fmt.Println("There is a wild duck.")
	wild := WildDuck{}
	wild.fly()
	wild.swim()

	// 声明一只家鸭
	fmt.Println("There is a domestic duck.")
	domestic := DomesticDuck{}
	domestic.swim()
}
