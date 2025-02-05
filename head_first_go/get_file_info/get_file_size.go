package main

import (
	"fmt"
	"log"
	"os"
)
// pdf P149
// go run get_file_size.go
func main() {
	fileInfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())  // 7
}