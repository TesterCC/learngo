package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// compare with fastapi demo, fastapi_demo/fastapi_tutorial/official_websocket_server_v2.py

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有CORS请求
	},
}

func getCookieOrToken(c *gin.Context) (string, error) {
	// 尝试从Cookie中获取session
	session, _ := c.Cookie("session")
	// 尝试从URL查询参数中获取token
	token := c.Query("token")
	fmt.Printf("[D] Sesison: %s\n", session)
	fmt.Printf("[D] Token: %s\n", token)

	if session == "" && token == "" {
		return "", &websocket.CloseError{Code: websocket.ClosePolicyViolation}
	}

	if session != "" {
		return session, nil
	}
	return token, nil
}

func websocketHandler(c *gin.Context) {
	itemID := c.Param("item_id")
	q := c.Query("q")

	fmt.Printf("[D] Item ID : %s\n", itemID)

	cookieOrToken, err := getCookieOrToken(c)
	if err != nil {
		log.Println("Error getting cookie or token:", err)
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to websocket:", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		responseMsg := "Session cookie or query token value is: " + cookieOrToken
		conn.WriteMessage(websocket.TextMessage, []byte(responseMsg))

		if q != "" {
			responseQueryMsg := "Query parameter q is: " + q
			conn.WriteMessage(websocket.TextMessage, []byte(responseQueryMsg))
		}

		responseDataMsg := "Message text was: " + string(message) + ", for item ID: " + itemID
		conn.WriteMessage(websocket.TextMessage, []byte(responseDataMsg))
	}
}

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./templates/index2.html")

	// 处理根路径的请求，返回index2.html
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index2.html", nil)
	})

	r.GET("/items/:item_id/ws", websocketHandler)
	r.Run("0.0.0.0:8000")
}