package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

/*
install lib: go get golang.ngrok.com/ngrok
run:
NGROK_AUTHTOKEN=your_ngrok_token go run ngrokdemo.go
Ingress established at: https://cc51-153-101-64-87.ngrok-free.app
test:
curl https://cc51-153-101-64-87.ngrok-free.app
Hello from ngrok-go ,it works!
*/

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	listener, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}

	log.Println("Ingress established at:", listener.URL())

	return http.Serve(listener, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from ngrok-go ,it works!")
}
