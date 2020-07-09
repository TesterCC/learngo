package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

/*
https://www.bilibili.com/video/BV18Q4y1M7NV?p=54
go get -v golang.org/x/text
go get -v golang.org/x/net/html
*/
var tagList = []string{"全部","美搭","美妆","居家","美食","母婴","科技","运动","汽车","文化","萌宠","园艺","动漫","星座","摄影","珠宝","旅游","型男","其他"}


func main() {
	fmt.Println(tagList)
	resp, err := http.Get("http://www.darenji.com/rank.html")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", resp.StatusCode)
		return
	}

	// 先判断网页编码
	e := determineEncoding(resp.Body)
	// GBK解码
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	//all, err := ioutil.ReadAll(resp.Body)
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", all)
}

func determineEncoding (r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)   // 因为r直接读了之后无法再读

	if err != nil {
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
