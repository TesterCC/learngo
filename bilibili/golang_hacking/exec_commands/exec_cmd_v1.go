package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// ref: https://www.bilibili.com/video/BV1b84y1z7HW/?p=4  Golang Hacking | 基于GO语言黑帽子 渗透测试

func executeCommand(command string, args_arr []string) (err error) {
	args := args_arr
	cmd_obj := exec.Command(command, args...)
	cmd_obj.Stdout = os.Stdout
	cmd_obj.Stderr = os.Stderr

	err = cmd_obj.Run()

	if err != nil {
		log.Fatal(err)
		return
	}
	return nil
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("get current directory failed:", err)
		return
	}
	fmt.Printf("current path: %s\n", dir)

	// normal level
	//cmd := "date"
	cmd := "ls"
	// ls -al
	executeCommand(cmd, []string{"-al"})

}
