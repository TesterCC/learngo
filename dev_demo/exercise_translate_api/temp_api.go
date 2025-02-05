package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type MarkedInfo struct {
	TaskId int32  `json:"task_id" `
	Header string `json:"header" `
	Body   string `json:"body" `
}

func SendMarkedInfo(c *gin.Context) {
	var markedInfo MarkedInfo
	if err := c.ShouldBindJSON(&markedInfo); err == nil {
		// 默认无缩进格式
		//data, err := json.Marshal(initConfig)
		// 带JSON缩进格式写文件
		data, err := json.MarshalIndent(markedInfo, "", "    ")
		if err != nil {
			fmt.Printf("json.Marshal failed,err:", err)
			//log.Fatalf("json.Marshal failed,err: ", err)
			return
		}

		fmt.Printf("%s\n", string(data))

		if data == nil {
			return
		}

		// write data in file, just test
		file, _ := os.Create("./dev_demo/exercise_translate_api/marked_info.json")
		defer file.Close()
		file.Write(data)

		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "request success", "data": markedInfo})

	} else {
		fmt.Printf("json.Marshal failed, err: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Server Internal Error", "data": nil})
	}
}

// simple demo, test on http://127.0.0.1:8888/send_raw_marked_info
func main() {
	// ReleaseMode DebugMode TestMode, default mode is DebugMode
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/send_raw_marked_info", SendMarkedInfo)

	// run server
	r.Run(":8888")

}
