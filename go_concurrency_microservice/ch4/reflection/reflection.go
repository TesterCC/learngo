package main

import (
	"fmt"
	"reflect"
)

// P74 定义一个简单结构体用于后续实验

// 定义 Person接口
type Person interface {
	SayHello(name string)
    Run() string
}

type Hero struct {
	Name string
	Age int
	Speed int
}

func (hero *Hero) SayHello(name string){
	fmt.Println("Hello " + name, ", I am " + hero.Name)
}

func (hero *Hero) Run() string {
	fmt.Println("I am running at speed ", hero.Speed)
	return "Running"
}

func main() {
	// 获取实例的反射 类型对象
	typeOfHero := reflect.TypeOf(Hero{})
	fmt.Printf("Hero's type is %s, kind is %s\n",typeOfHero, typeOfHero.Kind())
	fmt.Printf("Hero's type is %s, kind is %s\n",reflect.TypeOf(&Hero{}), reflect.TypeOf(&Hero{}).Kind())

	typeOfHeroElement := reflect.TypeOf(&Hero{}).Elem()
	fmt.Printf("typeOfHero elem : %s\n", typeOfHeroElement)
	
}
