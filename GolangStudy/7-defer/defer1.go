package main

import "fmt"

/*
知识点一： defer的执行顺序
*/

func func1()  {
	fmt.Println("A")
}

func func2()  {
	fmt.Println("B")
}

func func3()  {
	fmt.Println("C")
}

func main() {
	defer func1()
	defer func2()
	defer func3()
	// C B A
}
