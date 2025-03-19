package main

import (
	"log"
	"os"
	"os/exec"
)

// ref: https://www.bilibili.com/video/BV1b84y1z7HW/?p=4  Golang Hacking | 基于GO语言黑帽子 渗透测试

//func executeCommand(command string, args_arr []string) (err error) {
//	args := args_arr
//
//
//}

func main() {
	// basic
	cmd := "date"

	cmd_obj := exec.Command(cmd)
	cmd_obj.Stdout = os.Stdout
	cmd_obj.Stderr = os.Stderr

	err := cmd_obj.Run()

	if err != nil {
		log.Fatal(err)
	}

}
