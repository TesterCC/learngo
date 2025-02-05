package main

import (
	"log"
	"os"
)

// P62

func main() {
	err := os.Remove("222.txt")
	if err != nil{
		log.Fatal(err)
	}
}
