package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// CBCEncrypt 使用CBC模式加密
func CBCEncrypt(key, plaintext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plaintext = pkcs7Padding(plaintext, blockSize)
	ciphertext := make([]byte, blockSize+len(plaintext))
	iv := ciphertext[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blockSize:], plaintext)
	return hex.EncodeToString(ciphertext), nil
}

// CBCDecrypt 使用CBC模式解密
func CBCDecrypt(key []byte, encryptedHex string) ([]byte, error) {
	ciphertext, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(ciphertext) < blockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:blockSize]
	ciphertext = ciphertext[blockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	return pkcs7UnPadding(ciphertext), nil
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

	/*
		CBC（密码块链接模式，Cipher Block Chaining）：
		在加密前，每个明文块会和前一个密文块进行异或操作，第一个明文块则和 IV 进行异或。
		这种模式能保证相同的明文块加密后得到不同的密文块，增强了安全性。
	*/

	// CBC模式
	encryptedCBC, err := CBCEncrypt(key, text)
	if err != nil {
		fmt.Println("CBC Encryption failed:", err)
		return
	}
	fmt.Println("CBC Encrypted:", encryptedCBC)
	decryptedCBC, err := CBCDecrypt(key, encryptedCBC)
	if err != nil {
		fmt.Println("CBC Decryption failed:", err)
		return
	}
	fmt.Println("CBC Decrypted:", string(decryptedCBC))
}
