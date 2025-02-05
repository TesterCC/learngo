package main

import "fmt"

// P48

const name = "Golang is Good!"
const (
	surname = "王"
	personalName = "小明"
)

type aliasInt = int // 定一个一个类型别名
type myInt aliasInt // 定义一个新的类型

func main() {
	fmt.Println(name, surname, personalName)

	var alias aliasInt
	fmt.Printf("alias value is %v, type is %T\n", alias, alias)

	var myint myInt
	fmt.Printf("myint value is %v, type is %T\n", myint, myint)
	
}
