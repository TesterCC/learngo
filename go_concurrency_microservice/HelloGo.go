package main

// P38-39 基于图灵API实现一个简单的聊天机器人

// 请求体结构
type requestBody struct {
	Key string `json:"key"`
	Info string `json:"info"`
	UserId string `json:"userid"`
}

// 结果体结构
type responseBody struct {
	Code int `json:"code"`
	Text string `json:"text"`
	List []string `json:"list"`
	Url string `json:"url"`
}

// 请求机器人
func process(inputChan <- chan string, userid string){
	for {
		// 从通道中接受输入
		input := <- inputChan
		if input == "EOF"{
			break
		}
		// 构建请求体 todo
	}
}

func main() {
	
}
