package main

import "fmt"

//P14

func normal(x int) {
	if x == 1 {
		fmt.Println("X is equal to 1")
	} else {
		fmt.Println("X is not equal to 1")
	}
}

func switchx(x string) {
	switch x {
	case "java":
		fmt.Println("Use Java")
	case "python":
		fmt.Println("Use Python")
	default:
		fmt.Println("Use Go")
	}
}

func main() {
	//var x int
	var x = 2
	normal(x)

	switchx("go")
	switchx("python")

}
