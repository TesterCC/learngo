package main

import "fmt"

// P67-68   如果内嵌的类型为结构体，就可以直接访问内嵌结构体中的所有成员
type Wheel struct {
	shape string
}

type Car struct {
	Wheel
	Name string

}

func main() {
	car := &Car{
		Wheel{
			"Circle",
		},
		"Ford",
	}

	fmt.Println(car.Name, car.shape)
}
