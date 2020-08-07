package main

import "fmt"

// P65 为Person结构体添加修改姓名和输出个人信息两个方法

type Person struct {
	Name string   // 姓名
	Birth string  // 生日
	ID int64   // 身份证号
}

// 指针类型接收器，修改个人信息
func (person *Person) changeName(name string) {
	person.Name = name
}

// 非指针类型接收器，打印个人信息
func (person Person) printMess(){
	fmt.Printf("My name is %v, and my birthday is %v, and my id is %v\n",
		person.Name, person.Birth, person.ID)
	// 尝试修改个人ID，但对原接收器并没有影响。 因为非指针类型的接收器传递的是方法调用时接收器的一份值拷贝，并不会影响到原接收器。
	person.ID = 3
}


func main() {
	p1 := Person{
		Name: "Alice",
		Birth: "1990-01-02",
		ID: 1,
	}

	p1.printMess()

	p1.changeName("Bob")

	p1.printMess()
}
