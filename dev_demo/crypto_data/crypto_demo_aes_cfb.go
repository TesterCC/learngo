package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

/*
CFB（密码反馈模式，Cipher Feedback，）：是一种流加密模式，它将分组密码转换为流密码来处理数据。
CFB 模式把 AES 分组密码转换为流密码，这意味着它可以按字节（或位）对数据进行加密，而无需将数据分割成固定大小的块。
这种特性使得 CFB 模式能够处理任意长度的数据，即使数据长度不是 AES 块大小（通常为 16 字节）的整数倍，也无需进行填充操作。
*/
// encrypt 使用AES-CFB模式加密数据
// 参数:
//   - key: 16字节的密钥
//   - text: 要加密的明文数据
//
// 返回:
//   - 十六进制编码的密文字符串
//   - 错误信息(如果有)
func encrypt(key, text []byte) (string, error) {
	// 创建AES密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 创建密文缓冲区，大小为IV长度(16字节)加上明文长度
	ciphertext := make([]byte, aes.BlockSize+len(text))

	// 提取IV部分
	iv := ciphertext[:aes.BlockSize]

	// 生成随机IV
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 创建CFB加密器
	stream := cipher.NewCFBEncrypter(block, iv)

	// 执行加密操作
	stream.XORKeyStream(ciphertext[aes.BlockSize:], text)

	// 将密文转换为十六进制字符串
	return hex.EncodeToString(ciphertext), nil
}

// decrypt 使用AES-CFB模式解密数据
// 参数:
//   - key: 16字节的密钥
//   - encryptedHex: 十六进制编码的密文字符串
//
// 返回:
//   - 解密后的明文
//   - 错误信息(如果有)
func decrypt(key []byte, encryptedHex string) ([]byte, error) {
	// 将十六进制字符串解码为字节数组
	ciphertext, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return nil, err
	}

	// 检查密文长度是否足够
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("密文长度不足")
	}

	// 创建AES密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 提取IV
	iv := ciphertext[:aes.BlockSize]

	// 创建CFB解密器
	stream := cipher.NewCFBDecrypter(block, iv)

	// 准备明文缓冲区
	plaintext := make([]byte, len(ciphertext)-aes.BlockSize)

	// 执行解密操作
	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])

	return plaintext, nil
}

func main() {
	// 创建16字节的密钥
	key := []byte("example key 1234") // key need 16
	//key := []byte("test key 1234")  // key 13 will cause report

	fmt.Println("key length: ", len(key)) // key need 16, debug

	// 要加密的明文
	text := []byte("Hello World by Go!")

	// 加密
	encrypted, err := encrypt(key, text)
	if err != nil {
		fmt.Println("Encryption failed:", err)
		return
	}
	fmt.Println("Encrypted:", encrypted)

	// 解密
	decrypted, err := decrypt(key, encrypted)
	if err != nil {
		fmt.Println("Decryption failed:", err)
		return
	}
	fmt.Println("Decrypted:", string(decrypted))
}
