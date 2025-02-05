package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

/*
AES加密和解密
Advanced Encryption Standard, AES 又名 Rijndael 是 NIST 于 2001 年创建的一种加密算法。它使用 128 位数据块进行加密，是一种对称块密码。
ref: https://mp.weixin.qq.com/s?__biz=MzU2ODc0NTM1NQ==&mid=2247486484&idx=1&sn=0fa952b531e780917a79e3885460f574&chksm=fc8800f0cbff89e6d5904e61d4f92007ec24b48150fd34559cee8b492ed68cf09c79983509df&scene=132#wechat_redirect
 */
func main() {

	// cipher key
	key := "thisis32bitlongpassphraseimusing"

	// plaintext
	pt := "This is a secret"

	// send key to EncryptAES to encrypt plaintext
	c := EncryptAES([]byte(key), pt)

	// plaintext
	fmt.Println(pt)

	// ciphertext
	fmt.Println(c)

	// decrypt
	DecryptAES([]byte(key), c)
}

func EncryptAES(key []byte, plaintext string) string {

	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	fmt.Println("[+] DECRYPTED:", s)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}