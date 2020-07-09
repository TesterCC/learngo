package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])   // 因为os.Args[0]是命令本身的名字

	for i,ch := range os.Args {
		fmt.Printf("%d -> %s\n",i,ch)
	}
}
