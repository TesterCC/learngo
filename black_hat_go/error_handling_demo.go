package main

import (
	"errors"
	"fmt"
)

//P17-18

//type error interface {
//	Error() string
//}

// a custom error you can use, a user-defined string type
type MyError string

func (e MyError) Error() string {
	return string(e)
}

func foo2() error {
	return errors.New("Some Error Occurred")
}

func main() {
	if err := foo2(); err != nil {
		//Handle the error
		fmt.Println(err)    // print error info
	}
}
