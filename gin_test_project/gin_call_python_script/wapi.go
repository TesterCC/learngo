package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
)

type RequestParams struct {
	TaskId  string `json:"task_id"`
	Dir  string `json:"dir"`
	Path  string `json:"path"`
}

// because the api is simple, use gin directly
/*
go cross compiling
ref: https://go.dev/doc/install/source#environment
ref: https://www.jianshu.com/p/3ff21f818182


默认go build包含debug信息和符号表，减少生成二进制文件大小约30%的方法
go build -ldflags "-w -s" xxx.go

windows cross compile linux elf, default is windows exe
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build wapi.go

减少生成二进制文件大小约30%的方法
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" wapi.go

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
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Cors())


	r.POST("/executeCmd", func(c *gin.Context) {
		var params RequestParams
		if err := c.BindJSON(&params); err!= nil {
			c.JSON(400, gin.H{"error": "Invalid JSON data"})
			return
		}

		// 执行命令
		cmd := exec.Command("python", "cli.py", "-i", params.TaskId, "-d", params.Dir, "-f", params.Path)
		output, err := cmd.CombinedOutput()
		fmt.Println("[D] execute cmd: ", string(output)) // only debug use, bytes to string use string()
		if err!= nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// 处理输出为 JSON
		var data interface{}
		err = json.Unmarshal(output, &data)
		if err!= nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, data)
	})

	r.Run(":7777")
}
