package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 用gpt辅助完成 a basic websocket chat room
// compare with fastapi demo, fastapi_demo/fastapi_tutorial/official_websocket_server_v3.py

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有CORS请求
	},
}

type ConnectionManager struct {
	activeConnections map[*websocket.Conn]bool
	mu                sync.Mutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		activeConnections: make(map[*websocket.Conn]bool),
	}
}

//func (manager *ConnectionManager) connect(conn *websocket.Conn) {
//	manager.mu.Lock()
//	manager.activeConnections[conn] = true
//	manager.mu.Unlock()
//}

func (manager *ConnectionManager) connect(conn *websocket.Conn) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	if _, exists := manager.activeConnections[conn]; !exists {
		log.Println("New connection added.")
		manager.activeConnections[conn] = true
	}
}


//func (manager *ConnectionManager) disconnect(conn *websocket.Conn) {
//	manager.mu.Lock()
//	delete(manager.activeConnections, conn)
//	manager.mu.Unlock()
//	conn.Close()
//}

func (manager *ConnectionManager) disconnect(conn *websocket.Conn) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	if _, exists := manager.activeConnections[conn]; exists {
		log.Println("Connection disconnected.")
		delete(manager.activeConnections, conn)
		conn.Close()
	}
}

//func (manager *ConnectionManager) sendPersonalMessage(message string, conn *websocket.Conn) error {
//	return conn.WriteMessage(websocket.TextMessage, []byte(message))
//}
//
//func (manager *ConnectionManager) broadcast(message string) {
//	manager.mu.Lock()
//	defer manager.mu.Unlock()
//	for conn := range manager.activeConnections {
//		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
//			log.Printf("broadcast error: %v", err)
//			manager.disconnect(conn)
//		}
//	}
//}

func (manager *ConnectionManager) sendPersonalMessage(message string, conn *websocket.Conn) error {
	if _, exists := manager.activeConnections[conn]; exists {
		return conn.WriteMessage(websocket.TextMessage, []byte(message))
	}
	return nil // 或适当的错误处理
}

func (manager *ConnectionManager) broadcast(message string) {
	manager.mu.Lock()
	defer manager.mu.Unlock()
	for conn := range manager.activeConnections {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Printf("Broadcast error: %v, disconnecting...", err)
			delete(manager.activeConnections, conn)
			conn.Close()
		}
	}
}


func main() {
	r := gin.Default()
	manager := NewConnectionManager()

	r.LoadHTMLFiles("./templates/index3.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index3.html", nil)

		// 如果支持是html:='''<html>xxx</html>''', 用如下方法：
		//c.Writer.WriteHeader(http.StatusOK)
		//c.Writer.Write([]byte(html)) // html变量包含上面给出的HTML代码
	})

	r.GET("/ws/:client_id", func(c *gin.Context) {
		clientID := c.Param("client_id")
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer manager.disconnect(conn)

		manager.connect(conn)
		manager.broadcast(fmt.Sprintf("Client #%s joined the chat", clientID))

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Error: %v", err)
				}
				break
			}
			manager.sendPersonalMessage(fmt.Sprintf("You wrote: %s", string(message)), conn)
			manager.broadcast(fmt.Sprintf("Client #%s says: %s", clientID, string(message)))
		}

		manager.broadcast(fmt.Sprintf("Client #%s left the chat", clientID))
	})

	r.Run("0.0.0.0:8000")
}
