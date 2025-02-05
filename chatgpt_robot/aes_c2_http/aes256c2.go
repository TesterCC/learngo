
/*
用Go语言和Gin框架实现http c2服务器，且支持AES-256加密通信
在这个示例中，使用了Gin框架来实现HTTP C2服务器的接口，
通过/ping接口实现心跳检测，
通过/command接口接收客户端发送的指令，并对指令使用AES-256加密算法进行解密。

 */

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	key = "abcdefghijklmnopqrstuvwxyzabcdef"   // your key
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.POST("/command", func(c *gin.Context) {
		decryptedData, err := decryptAES(c.PostForm("data"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 这里可以处理从客户端发送过来的指令
		fmt.Println("received command:", string(decryptedData))

		c.JSON(http.StatusOK, gin.H{"message": "command received"})
	})
	r.Run()
}

func decryptAES(ciphertext string) ([]byte, error) {
	ciphertextBytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertextBytes, ciphertextBytes)

	return ciphertextBytes, nil
}
