package main

import "fmt"

// P12
type Person struct {
	Name string
	Age  int
}

type Dog struct {
}

// P13 struct and interface
type Friend interface {
	SayHello()
}

// method
func (d *Dog) SayHello() {
	fmt.Println("Woof woof")
}

func Greet(f Friend) {
	f.SayHello()
}

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
