package main

import (
	"log"
	"os"
)

// P61-62

func main() {
	newFile, err := os.Create("000.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()

	originPath := "000.txt"
	newPath := "222.txt"

	// 短变量声明左侧变量名不能重复
	err2 := os.Rename(originPath, newPath)

	if err2 != nil {
		log.Fatal(err2)
	}
}
