package main

import (
	"net/http"
)


// function is just a normal function; it doesn’t have a ServeHTTP() method. So in itself it isn’t a handler.
// ref: https://lets-go.alexedwards.net/sample/02.09-the-http-handler-interface.html
func handler (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's works...\n"))   // normal
	//io.WriteString(w, "It's works...\n")
}



func main() {

    http.HandleFunc("/", handler)
	http.ListenAndServe(":7777", nil)
}
