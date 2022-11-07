package main

import (
	"fmt"
	"os"
	"os/exec"
)

/*
use exec package to run shell command
ref:
https://mp.weixin.qq.com/s/yitQcfP9KEmwvViL-1K5wA
*/

// run ls
func main() {
	//cmd := exec.Command("ls")   // just run system command
	cmd := exec.Command("ls", "-al")   // just run system command

    cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	e := cmd.Run()
	CheckError(e)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
