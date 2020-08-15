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
	Name  string
	Age   int
	Speed int
}

func (hero *Hero) SayHello(name string) {
	fmt.Println("Hello "+name, ", I am "+hero.Name)
}

func (hero *Hero) Run() string {
	fmt.Println("I am running at speed ", hero.Speed)
	return "Running"
}

func main() {
	// 获取实例的反射 类型对象
	//typeOfHero := reflect.TypeOf(Hero{})
	//fmt.Printf("Hero's type is %s, kind is %s\n",typeOfHero, typeOfHero.Kind())
	//fmt.Printf("Hero's type is %s, kind is %s\n",reflect.TypeOf(&Hero{}), reflect.TypeOf(&Hero{}).Kind())
	//
	//typeOfHeroElement := reflect.TypeOf(&Hero{}).Elem()
	//fmt.Printf("typeOfHero elem : %s\n", typeOfHeroElement)

	typeOfHero := reflect.TypeOf(Hero{})

	// 通过 NumField 获取结构体字段的数量
	for i := 0; i < typeOfHero.NumField(); i++ {
		fmt.Printf("field name is %s, type is %s, kind is %s\n",
			typeOfHero.Field(i).Name,
			typeOfHero.Field(i).Type,
			typeOfHero.Field(i).Type.Kind())
	}

	// 获取名称为 name 的成员字段类型对象
	nameField, _ := typeOfHero.FieldByName("Name")
	fmt.Printf("field name is %s, type is %s, kind is %s\n", nameField.Name, nameField.Type, nameField.Type.Kind())

	// P78-82 todo
}
