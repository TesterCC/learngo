package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

// dev 20240806 add web api
type RequestAPIParams struct {
	TaskId     string `json:"i"`
	Dir        string `json:"d"`
	File       string `json:"f"`
	Statistics bool   `json:"s"`
}

type JsonRespStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}


func ReturnResp(c *gin.Context, code int, msg interface{}, data interface{}) {
	jsonData := &JsonRespStruct{Code: code, Msg: msg, Data: data}
	c.JSON(http.StatusOK, jsonData)
}

// because the api is simple, use gin directly
/*
go cross compiling
ref: https://go.dev/doc/install/source#environment
ref: https://www.jianshu.com/p/3ff21f818182


默认go build包含debug信息和符号表，减少生成二进制文件大小约30%的方法
go build -ldflags "-w -s" xxx.go

windows cross compile linux elf, default is windows exe
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build gapi.go

减少生成二进制文件大小约30%的方法
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" gapi.go

*/

// custom middleware
// cross origin middleware

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()   // gin.Default() 会创建一个带有一些默认中间件（如 Logger 和 Recovery 中间件）的 Gin 引擎实例。
	//r := gin.New()     // gin.New() 创建的是一个不包含这些默认中间件的 Gin 引擎实例。
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())

	// custom middleware
	r.Use(Cors())

	r.POST("/StartCodeGene", func(c *gin.Context) {
		var params RequestAPIParams
		if err := c.BindJSON(&params); err != nil {

			//c.JSON(http.StatusBadRequest, gin.H{
			//	"code": http.StatusBadRequest,
			//	"msg": "json parse error, please check input data",
			//	"data": nil,
			//})
			ReturnResp(c, http.StatusBadRequest, "json parse error, please check input data", nil)
			return
		}

		fmt.Printf("[D] TaskId: %s, Dir: %s, File: %s, Statistics: %t\n", params.TaskId, params.Dir, params.File, params.Statistics)

		// first check
		if params.TaskId == "" {
			ReturnResp(c, http.StatusBadRequest, "task id is empty", nil)
			return
		}

		// run cmd
		//pyCompiler := "/opt/tools/apt_parser/apt-env/bin/python3"  // deprecated, debug use in dev vm venv
		pyCompiler := "/usr/bin/python3"   // release use

		var cmd *exec.Cmd
		if params.Statistics {
			// only get rule info use
			//cmd = exec.Command("/opt/tools/apt_parser/apt-env/bin/python3", "-W ignore", "/opt/tools/apt_parser/cli.py", "-s")
			cmd = exec.Command(pyCompiler, "-W ignore", "/opt/tools/apt_parser/cli.py", "-s")
		} else {
			if params.Dir != "" {
				//cmd = exec.Command("/opt/tools/apt_parser/apt-env/bin/python3", "-W ignore", "/opt/tools/apt_parser/cli.py", "-i", params.TaskId, "-d", params.Dir)
				cmd = exec.Command(pyCompiler, "-W ignore", "/opt/tools/apt_parser/cli.py", "-i", params.TaskId, "-d", params.Dir)
			} else if params.File != "" {
				//cmd = exec.Command("/opt/tools/apt_parser/apt-env/bin/python3", "-W ignore", "/opt/tools/apt_parser/cli.py", "-i", params.TaskId, "-f", params.File)
				cmd = exec.Command(pyCompiler, "-W ignore", "/opt/tools/apt_parser/cli.py", "-i", params.TaskId, "-f", params.File)
			} else {
				ReturnResp(c, http.StatusBadRequest, "some args empty", nil)
				return
			}
		}

		//cmd := exec.Command("python -W ignore", "cli.py", "-i", params.TaskId, "-d", params.Dir, "-f", params.Path)
		output, err := cmd.CombinedOutput()
		fmt.Println("[D] execute cmd: ", string(output)) // only debug use, bytes to string use string()
		if err != nil {
			ReturnResp(c, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		// output json
		var data JsonRespStruct
		err = json.Unmarshal(output, &data)
		if err != nil {
			//c.JSON(500, gin.H{"error": err.Error()})
			ReturnResp(c, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		//c.JSON(200, data)
		ReturnResp(c, 1,"success", data.Data)
		return
	})

	// add version info print for checking
	gapiVersion := "0.1.0.20240809"
	fmt.Println("[Dev] current gapi version: ", gapiVersion)
	r.Run("0.0.0.0:7777")
}
