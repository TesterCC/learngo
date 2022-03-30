package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

/*
需求：根据参数触发api执行系统命令

go mod init api_send_data

install:
go env -w GOPROXY=https://goproxy.cn
go get -u github.com/gin-gonic/gin

build:
build & run
go build -o api_send_data api_run_cmd.go
./api_send_data

testcase:
curl -X GET "http://127.0.0.1:65522/send_2022_testdata?first_arg={first_cus_arg}&last_arg={last_cus_arg}"
curl -X GET "http://127.0.0.1:65522/send_2022_testdata?first_arg=aaabbb&last_arg=20220330"
*/

var (
	first_cus_arg = "aaabbb"
	last_cus_arg = "20220330"
)


func ExecCommand(strCommand string)(string){
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


func main() {
	r := gin.Default()

	//获取GET请求参数
	r.GET("/send_2022_testdata", func(c *gin.Context) {
		// 获取一个不带默认值的参数
		firstArg := c.Query("first_arg")
		// 获取一个带默认值的参数
		lastArg := c.DefaultQuery("last_arg", "eee")

		if firstArg != first_cus_arg {
			return
		} else if lastArg != last_cus_arg {
			return
		} else {
			fd_time := time.Now().Format("2006-01-02 15:04:05")

            // 执行系统命令
            str_data := ExecCommand("nohup python3 /opt/test/sa/send_syslog_100eps/send_cycle.py > /dev/null 2>&1 &")

			c.String(http.StatusOK, "[GET] [%s] Launch send test data operations successfully.\n[+]%s", fd_time, str_data)
		}

	})

	r.Run(":65522")  //  指定端口启动
}
