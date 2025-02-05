package main

import (
	"fmt"
	"log"
	"net"
)

// 获取本机IP的两种方式
// get localhost ip
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _,addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback(){
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast(){
			continue
		}
		return ipAddr.IP.String(),nil
	}
	return
}


// Get preferred outbound ip of this machine
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println("[+] Detail info:",localAddr.String())
	return localAddr.IP.String()
}

func main() {
	ip, err := GetLocalIP()
	if err != nil {
		return
	}
	fmt.Println("[+] Local IP:", ip)

	GetOutboundIP()
}
