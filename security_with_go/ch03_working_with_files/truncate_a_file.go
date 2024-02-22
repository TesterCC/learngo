package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// security with go - P60
	// Truncate a file to 100 bytes.  ls -lht 100B
	// It will fill any blank space with null bytes.

	/// go time.Now().Unix() , python time.time()
	outputFileName := fmt.Sprintf("./test_%d.txt", time.Now().Unix())

	// malicious actors may truncate files just for the sake of destroying data
	err := os.Truncate(outputFileName, 100)

	if err != nil {
		log.Fatal(err)
	}
}
