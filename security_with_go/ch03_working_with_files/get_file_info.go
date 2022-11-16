package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// print out all the metadata available about a file
	// Stat returns file info. It will return an error if there is no file.
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name: ", fileInfo.Name())
	fmt.Println("Size in bytes: ", fileInfo.Size())
	fmt.Println("Last modified: ", fileInfo.Mode())
}