package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.141592))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, Go!"))
}
