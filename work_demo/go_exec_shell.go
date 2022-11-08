package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

/*
use exec package to run shell command
ref:
https://mp.weixin.qq.com/s/yitQcfP9KEmwvViL-1K5wA

go tool dist list   // display OS and ARCH composition
*/

// run ls
func main() {
	//cmd := exec.Command("ls")   // just run system command

	if runtime.GOOS == "darwin" {
		fmt.Println("Apple!")
		cmd := exec.Command("ls", "-al", "/tmp")   // just run system command
		// output terminal info
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		e := cmd.Run()
		CheckError(e)
	} else {
		fmt.Println(">>>")
		cmd := exec.Command("whoami")   // just run system command
		// output terminal info
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		e := cmd.Run()
		CheckError(e)
	}


}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
