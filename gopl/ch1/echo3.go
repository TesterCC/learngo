package main

import (
	"fmt"
	"os"
	"strings"
)

// P6 use join
// ./echo3 -v -t -h
func main() {
	fmt.Println(strings.Join(os.Args[1:]," "))   // 因为os.Args[0]是命令本身的名字
}
