package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱安全研发"
	fmt.Println(len(s))   // 21

	fmt.Printf("%X\n", []byte(s))  // %X print hex

	for _,b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i,ch := range s {  // ch is a rune，rune is alias of int32
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i,ch := range[]rune(s){
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
