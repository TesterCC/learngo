package main

import (
	"log"
	"os"
)

func main() {
	// security with go - P60
	// Truncate a file to 100 bytes.  ls -lht 100B
	// It will fill any blank space with null bytes.
	// malicious actors may truncate files just for the sake of destroying data

	err := os.Truncate("test.txt", 100)

	if err != nil {
		log.Fatal(err)
	}
}
