package main

import "net/http"

// https://www.bilibili.com/video/BV18Q4y1M7NV?p=53

func main() {
	resp, err := http.Get("http://www.darenji.com/rank.html")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

	}
}
