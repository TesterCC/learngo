package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/*
P12
1.5 获取一个url

fetch输出从URL获取的内容

run in terminal:
go run get_a_url.go http://www.xxx.com
go run get_a_url.go https://paper.seebug.org/
go run get_a_url.go https://paper.seebug.org/rss/
*/

func main() {
	for _,url := range os.Args[1:] {
		resp,err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
