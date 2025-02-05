package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// compare with fastapi demo, fastapi_demo/fastapi_tutorial/official_websocket_server.py

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins for simplicity in this example
		return true
	},
}

func main() {
	// 将Gin设置为发布模式
	//gin.SetMode(gin.DebugMode)

	r := gin.Default()

	// Load HTML templates
	// todo Attention: need run in right dir, such  /d/git_workspace/ws_go/xxx/gin_test_project/websocket
	r.LoadHTMLFiles("./templates/index.html")

	//// 静态文件服务，服务当前目录下的templates文件夹内容
	//r.Static("/static", "./static")


	// 处理根路径的请求，返回index.html
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// WebSocket路由
	r.GET("/ws", func(c *gin.Context) {
		conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			http.NotFound(c.Writer, c.Request)
			return
		}
		defer conn.Close()

		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				return // Automatically handles disconnection.
			}

			// 要在你的WebSocket服务器的响应中添加前缀 "Received data: "，你只需修改消息回复的字符串，将接收到的消息（p）前加上这个前缀。
			// 因为gorilla/websocket库的WriteMessage方法接收一个字节切片作为消息参数，你需要先将前缀和接收到的消息组合成一个新的字符串，然后将这个字符串转换为字节切片。
			// Echo the received message back to the client.
			// 组合消息，加上前缀 "Received data: "
			responseMessage := "Message text was: " + string(p)

			if err := conn.WriteMessage(messageType, []byte(responseMessage)); err != nil {
				return // Again, handles errors in an easy way.
			}
		}
	})

	// 运行服务器
	r.Run("0.0.0.0:8000")

}