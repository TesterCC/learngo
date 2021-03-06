package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

const LoginUrl = "https://10.0.0.115/api/login/"

// It's so inconvenient...
func main() {
	// 通过设置tls.Config的InsecureSkipVerify为true，client将不再对服务端的证书进行校验。
	tr := &http.Transport{
		TLSClientConfig:        &tls.Config{InsecureSkipVerify:true},
	}

	client := &http.Client{Transport:tr}

	// simulate a get post
	//resp, err := http.Get(LoginUrl)
	resp, err := client.Get(LoginUrl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // read body content
	fmt.Println(string(body))
}


