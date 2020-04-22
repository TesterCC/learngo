package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// Task: gin框架中结合gorilla实现webSocket
// gin框架写rest接口很方便且性能好
// ref: https://blog.csdn.net/lihao19910921/article/details/81907742

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket请求ping，返回pong

func ping(c *gin.Context) {
	// 升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	defer ws.Close()

	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()

		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}

}

func main() {
	bindAddress := "localhost:3030"
	r := gin.Default()
	r.GET("/ping", ping)  // 传入func ping
	r.Run(bindAddress)

}
