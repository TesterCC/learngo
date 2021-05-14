package main

import (
	"os"
	"strings"
	"time"
)

/* ref network_monitor.py
gopsutil

doc: https://pkg.go.dev/github.com/shirou/gopsutil#section-documentation

install：
go get github.com/shirou/gopsutil

for devops and suspicious process analysis
*/

func log_output(str_content string) {
	fd, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd_time := time.Now().Format("2006-01-02 15:04:05")
	fd_content := strings.Join([]string{"======", fd_time, "=====", str_content, "\n"}, "")
	buf := []byte(fd_content)
	fd.Write(buf)
	fd.Close()
}

func main() {
	log_output("test")
}
