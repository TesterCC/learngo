package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// check file exists
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func main() {

	// get file dir by file path
	resultFilePath := "/opt/tools/results/tid/result.json"
	directory := filepath.Dir(resultFilePath)
	fmt.Println("[D] result file path: ", directory)

	if fileExists(resultFilePath) {
		fmt.Println("[I] file is exist.")
	} else {
		fmt.Println("[I] file isn't exist.")
	}
}
