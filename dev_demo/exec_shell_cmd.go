package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

// https://blog.csdn.net/sz76211822/article/details/83270525
// Linux Go执行shell命令并得到返回结果
func ExecCommand(strCommand string)(string){
	// run in linux
	cmd := exec.Command("/bin/bash", "-c", strCommand)


	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil{
		fmt.Println("Execute failed when Start:" + err.Error())
		return ""
	}

	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return ""
	}
	return string(out_bytes)
}


func main(){
	strData := ExecCommand("ifconfig")
	fmt.Println("[+] Execute finished:" + strData)
}
