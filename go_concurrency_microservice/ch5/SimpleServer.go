package main

import "net/http"

func main() {
	http.ListenAndServe("127.0.0.1:9999", nil)
}
