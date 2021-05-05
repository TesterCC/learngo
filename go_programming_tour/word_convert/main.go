package main

import (
	"learngo/go_programming_tour/word_convert/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Excute err: %v", err)
	}
}
