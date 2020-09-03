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


	// P78-82 获取接口Person中方法的方法类型对象
	// 声明一个Person接口，并用Hero作为接收器
	var person Person = &Hero{}
	// 获取接口Person的类型对象
	typeOfPerson := reflect.TypeOf(person)
	// 打印Person的方法类型和名称
	for i := 0; i<typeOfPerson.NumMethod();i++{
		fmt.Printf("method is %s, type is %s, kind is %s.\n",
			typeOfPerson.Method(i).Name,
			typeOfPerson.Method(i).Type,
			typeOfPerson.Method(i).Type.Kind())
	}
	method,_ := typeOfPerson.MethodByName("Run")
	fmt.Printf("method is %s, type is %s, kind is %s.\n", method.Name, method.Type, method.Type.Kind())

	// 可以使用reflect.New方法根据变量的Type对象创建一个相同类型的新变量，值以Value对象的形式返回
	typeOfHero2 := reflect.TypeOf(Hero{})
	heroValue := reflect.New(typeOfHero2)
	fmt.Printf("Hero's type is %s, kind is %s\n", heroValue.Type(), heroValue.Kind())
}
