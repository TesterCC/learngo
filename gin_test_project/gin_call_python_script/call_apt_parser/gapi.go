package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
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

//type TaskResultStruct struct {
//	TaskID       string         `json:"task_id"`
//	AnalysisInfo []string `json:"analysis_info"`
//	APTInfo      []string `json:"apt_info"`
//}

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

func readFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		//fmt.Println("[E] read file failed:", err)
		return nil, err
	}
	return data, nil
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func main() {
	r := gin.Default() // gin.Default() 会创建一个带有一些默认中间件（如 Logger 和 Recovery 中间件）的 Gin 引擎实例。
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
		}

		// run cmd
		//pyCompiler := "/opt/tools/apt_parser/apt-env/bin/python3"  // deprecated, debug use in dev vm venv
		pyCompiler := "/usr/bin/python3" // release use

		var cmd *exec.Cmd
		if params.Statistics {
			// only get rule info use
			//cmd = exec.Command("/opt/tools/apt_parser/apt-env/bin/python3", "-W ignore", "/opt/tools/apt_parser/cli.py", "-s")
			cmd = exec.Command(pyCompiler, "-W ignore", "/opt/tools/apt_parser/cli.py", "-s")
		} else {
			// example: nohup /usr/bin/python3 /opt/root_node/main.py > /opt/a.log 2>&1 &
			if params.Dir != "" {
				//cmd = exec.Command(pyCompiler, "-W ignore", "/opt/tools/apt_parser/cli.py", "-i", params.TaskId, "-d", params.Dir)
				// attention
				cmd = exec.Command(pyCompiler,"-W ignore", "/opt/tools/apt_parser/cli.py", "-i", params.TaskId, "-d", params.Dir)

				// cmd.Start() 用于启动一个命令并立即返回，命令在后台异步执行，不会阻塞当前程序的执行
				AsyncErr := cmd.Start()
				if AsyncErr != nil {
					ReturnResp(c, http.StatusInternalServerError, "run time error, please check input cmd", nil)
				}
				ReturnResp(c, 1, "success", nil)
				return

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
		ReturnResp(c, 1, "success", data.Data)
		return
	})

	r.GET("/CheckRule", func(c *gin.Context) {
		pyCompiler := "/usr/bin/python3" // release use

		var cmd *exec.Cmd

		// just check rule
		cmd = exec.Command(pyCompiler, "-W ignore", "/opt/tools/apt_parser/cli.py", "-c")

		output, err := cmd.CombinedOutput()
		fmt.Println("[D] execute cmd: ", string(output)) // only debug use, bytes to string use string()

		// output json
		var data JsonRespStruct
		err = json.Unmarshal(output, &data)
		if err != nil {
			//c.JSON(500, gin.H{"error": err.Error()})
			ReturnResp(c, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		ReturnResp(c, 1, "success", data.Data)
		return
	})

	r.GET("/TaskResult", func(c *gin.Context) {

		//id := c.Param("id")  // "/find/:id"
		taskId := c.Query("ID") // /find?id=xxx
		fmt.Println("[D] input task id is: ", taskId)
		//fmt.Printf("[D] args type: %T\n ", taskId)  // string

		taskResultPath := fmt.Sprintf("/opt/tools/results/%s/result.json", taskId)  // release use it
		//taskResultPath := fmt.Sprintf("./results/%s/result.json", taskId) // local dev use
		fmt.Printf("[D] task result file path: %s\n", taskResultPath)

		// 1. check file exist
		if !fileExists(taskResultPath) {
			ReturnResp(c, http.StatusInternalServerError, "file does not exist", nil)
			return
		}

		// 2. read json
		output, err := readFile(taskResultPath)
		//fmt.Printf(">>>>>: %T\n", output)   // debug  []uint8
		fmt.Println("[D] read json content:\n", string(output)) // debug print

		if err != nil {
			fmt.Println("[E] json read error:", err)
			ReturnResp(c, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		// attention: 直接将 JSON 数据解析到一个 map[string]interface{} 中，不需要事先定义结构体
		var data map[string]interface{}
		err2 := json.Unmarshal(output, &data)

		// // test pass, can get data
		//value, ok := data["analysis_info"]
		//if ok {
		//	fmt.Println(">>> get value: ", value)
		//}
		// // test end

		if err2 != nil {
			fmt.Println("[E] json parse error:", err)
			ReturnResp(c, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		ReturnResp(c, 1, "success", data)
		return
	})

	// add version info print for checking
	gapiVersion := "0.2.1.20240816"
	gapiListenIP := "0.0.0.0"
	gapiListenPort := 7777

	launchServer := fmt.Sprintf("%s:%d", gapiListenIP, gapiListenPort)

	fmt.Println("[DEV] current gapi version: ", gapiVersion)
	fmt.Println("[DEV] current gapi server: ", launchServer)

	r.Run(launchServer)
}
