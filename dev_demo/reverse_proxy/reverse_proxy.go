package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

/*
test:
curl http://localhost:7070/ping
*/
func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		target, _ := url.Parse("https://blog.aictf.com/logo.svg")
		rp := httputil.NewSingleHostReverseProxy(target)
		// change input
		rp.Director = func(req *http.Request) {
			req.URL.Path = target.Path
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.Host = target.Host
		}
		rp.ServeHTTP(w, r)
	})
	http.ListenAndServe(":7070", nil)
}
