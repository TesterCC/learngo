package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

// P38-39 基于图灵API实现一个简单的聊天机器人  官方已经升级API到v2了

// 代表请求体和响应体的JSON格式
// 请求体结构
type requestBody struct {
	Key    string `json:"key"`
	Info   string `json:"info"`
	UserId string `json:"userid"`
}

// 结果体结构
type responseBody struct {
	Code int      `json:"code"`
	Text string   `json:"text"`
	List []string `json:"list"`
	Url  string   `json:"url"`
}

// 请求机器人
func process(inputChan <-chan string, userid string) {
	apiKey := "" // your api key

	for {
		// 从通道中接受输入
		input := <-inputChan
		if input == "EOF" {
			break
		}
		// 构建请求体
		reqData := &requestBody{
			Key:    apiKey,
			Info:   input,
			UserId: userid,
		}

		// 转义为json
		byteData, _ := json.Marshal(&reqData)

		// 请求聊天机器人接口
		req, err := http.NewRequest("POST", "http://tuling123.com/openapi/api", bytes.NewReader(byteData))
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Network Error!")
		} else {
			// 将结果从json中解析并输出到命令行
			body, _ := ioutil.ReadAll(resp.Body)
			var respData responseBody
			json.Unmarshal(body, &respData)
			fmt.Println("AI: " + respData.Text)
		}

		resp.Body.Close()

	}
}

func main() {
	var input string
	fmt.Println("Enter 'EOF' to shut down: ")
	// 创建通道
	channel := make(chan string)

	// main结束时关闭通道
	defer close(channel)

	// 启动goroutine运行机器人回答线程
	go process(channel, string(rand.Int63()))
	for {
		// 从命令行中读取输入
		fmt.Scanf("%s", &input)

		// 将输入放入到管道中
		channel <- input
		// 结束程序
		if input == "EOF" {
			fmt.Println("Bye!")
			break
		}

	}


}
