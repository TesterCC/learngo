//使用 Go 语言实现的 AES-256 加密通信的 TCP 客户端和服务端的示例

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
)

func encrypt(key []byte, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], text)
	return ciphertext, nil
}

func decrypt(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext is too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	key := []byte("examplekey16bytes")
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Println("error reading data:", err)
			return
		}
		plaintext, err := decrypt(key, buf[:n])
		if err != nil {
			log.Println("error decrypting data:", err)
			return
		}
		fmt.Println("Received message:", string(plaintext))
		ciphertext, err := encrypt(key, plaintext)
		if err != nil {
			log.Println("error encrypting data:", err)
			return
		}
		_, err = conn.Write(ciphertext)
		if err != nil {
			log.Println("error writing data:", err)
			return
		}
	}
}


func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}