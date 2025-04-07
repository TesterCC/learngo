package main

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

/*
ECB（电子密码本模式，Electronic Codebook）：
这是最简单的加密模式，它把明文分成固定大小的块，然后对每个块独立进行加密。
不过，相同的明文块会被加密成相同的密文块，所以这种模式缺乏安全性，不适合加密长文本。
*/

// ECBEncrypt 使用ECB模式加密
func ECBEncrypt(key, plaintext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plaintext = pkcs7Padding(plaintext, blockSize)
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += blockSize {
		block.Encrypt(ciphertext[i:i+blockSize], plaintext[i:i+blockSize])
	}
	return hex.EncodeToString(ciphertext), nil
}

// ECB decrypt 使用ECB模式解密
func ECBDecrypt(key []byte, encryptedHex string) ([]byte, error) {
	ciphertext, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += blockSize {
		block.Decrypt(plaintext[i:i+blockSize], ciphertext[i:i+blockSize])
	}
	return pkcs7UnPadding(plaintext), nil
}

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// pkcs7UnPadding 去除填充
func pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func main() {
	key := []byte("example key 1234")
	text := []byte("Hello World by Go!")

	// ECB模式
	encryptedECB, err := ECBEncrypt(key, text)
	if err != nil {
		fmt.Println("ECB Encryption failed:", err)
		return
	}
	fmt.Println("ECB Encrypted:", encryptedECB)
	decryptedECB, err := ECBDecrypt(key, encryptedECB)
	if err != nil {
		fmt.Println("ECB Decryption failed:", err)
		return
	}
	fmt.Println("ECB Decrypted:", string(decryptedECB))
}
