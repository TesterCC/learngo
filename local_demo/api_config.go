package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

// go get -u github.com/gin-gonic/gin

const configCenterConfig = "./config.json"

type config struct {
	Probing      string `json:"probing" `
	ProbingNum   string `json:"probingNum" `
	ControlNum   string `json:"controlNum" `
	Identify     string `json:"identify" `
	ControlIp    string `json:"controlIp" `
	Tactics      string `json:"tactics" `
	Cm           string `json:"cm" `
	Exploit      string `json:"exploit" `
	Pm           string `json:"pm" `
	Verification string `json:"verification" `
	Situation    string `json:"situation" `
}

//PathExists 判断一个文件或文件夹是否存在
//输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func configDetail(c *gin.Context) {
	// check config.json, if not create it; if existed, read it
	ret, err := PathExists(configCenterConfig)
	if err != nil {
		log.Fatalln(err)
	}

	// according check config.json result
	if ret == true {
		//fmt.Println("file exist, read it")

		filePtr, err := os.Open("config.json")
		if err != nil {
			fmt.Println("Open file failed [Err:%s]", err.Error())
			return
		}
		defer filePtr.Close()

		readConfig := config{}
		decoder := json.NewDecoder(filePtr)
		err = decoder.Decode(&readConfig)
		if err != nil {
			log.Printf("decode error [ %v ]", err)
			return
		} else {
			//fmt.Println(readConfig)
			//fmt.Println(readConfig.Cm)
			//fmt.Println(readConfig.ControlNum)
		}

		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "请求成功", "data": readConfig})

	} else {
		//fmt.Println("file not exist, need to create")
		initConfig := &config{}
		// 默认无缩进格式
		//data, err := json.Marshal(initConfig)
		// 带JSON缩进格式写文件
		data, err := json.MarshalIndent(initConfig, "", "    ")
		if err != nil {
			fmt.Printf("json.Marshal failed,err:", err)
			//log.Fatalf("json.Marshal failed,err: ", err)
			return
		}

		fmt.Printf("%s\n", string(data))

		file, _ := os.Create("config.json")
		defer file.Close()
		file.Write(data)

		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "请求成功", "data": nil})

	}
}

func configSync(c *gin.Context) {
	var syncConfig config
	if err := c.ShouldBindJSON(&syncConfig); err == nil {
		// 默认无缩进格式
		//data, err := json.Marshal(initConfig)
		// 带JSON缩进格式写文件
		data, err := json.MarshalIndent(syncConfig, "", "    ")
		if err != nil {
			fmt.Printf("json.Marshal failed,err:", err)
			//log.Fatalf("json.Marshal failed,err: ", err)
			return
		}

		fmt.Printf("%s\n", string(data))
		fmt.Printf("post args empty...")
		if data == nil {
			return
		}

		file, _ := os.Create("config.json")
		defer file.Close()
		file.Write(data)

		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "请求成功", "data": syncConfig})

	} else {
		fmt.Printf("json.Marshal failed, err: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "服务器内部错", "data": nil})
	}
}

// todo net request url
func configUpdate(c *gin.Context) {

}

// simple demo

func main() {
	// ReleaseMode DebugMode TestMode, default mode is DebugMode
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/config/detail", configDetail)
	r.POST("/config/sync", configSync)
	r.POST("/config/update", configUpdate)

	// run server
	r.Run(":9999")

}
