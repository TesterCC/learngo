package main

import "fmt"

// switch后可以没有表达式
func grade(score int) string{
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}


//普通的switch， go中的switch会自动break，除非使用fallthrough
func eval2(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:"+op)  //相当于报错，后面会详细讲
	}
	return result
}



func main() {
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(100),
		//grade(-3),
		)

	fmt.Println(
		eval2(3,4, "+"),
		eval2(7,6, "-"),
		eval2(5,6, "*"),
		eval2(6,3, "/"),
		//eval(6,3, "%"),
		)
}
