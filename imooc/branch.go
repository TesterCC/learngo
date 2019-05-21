package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"   // terminal can use relative path   // IDE need abs path
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("%s\n", contents)   // contents byte
	}
	
}