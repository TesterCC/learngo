package main

/*
use exec package to run shell command
ref:
https://mp.weixin.qq.com/s/yitQcfP9KEmwvViL-1K5wA

go tool dist list   // display OS and ARCH composition
*/


import (
"fmt"
"log"
"os/exec"
)

func main() {

	// get output
	out, err := exec.Command("ls", "-lht").Output()

	if err != nil {
		log.Fatal(err)
	}


	fmt.Println(string(out))
}