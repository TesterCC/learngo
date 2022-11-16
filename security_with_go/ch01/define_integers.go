package main

import "fmt"

// P34

func main() {
	// Decimal for 17
	number0 := 17

	// Octal for 17
	number1 := 017

	// Hexadecimal for 17
	number2 := 0X0F

	fmt.Println(number0, number1, number2)
	fmt.Printf("%T,%T,%T", number0, number1, number2)
}
