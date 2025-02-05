package main

import (
	"github.com/malfunkt/iprange"
	"log"
)

// 白帽子安全开放实战 P17-18

// 要实现对多IP的扫描，需要引入一个第三方包 "github.com/malfunkt/iprange"，它实现了类似于 Nmap风格对多个IP的解析。
// 实现对多IP的扫描，引入第三方包 github.com/malfunkt/iprange   安装：go get github.com/malfunkt/iprange
// iprange实现了对多个IP的解析

// iprange 示例
func main() {
	//list, err := iprange.ParseList("10.0.0.1, 10.0.0.5-10, 192.168.1.*,192.168.10.0/24")
	//list, err := iprange.ParseList("10.0.0.1, 10.0.0.5-10")

	//list, err := iprange.ParseList("192.168.1.*,192.168.10.0/24")

	// iprange库将Nmap风格的IP解析为 AddressRange对象，
	list, err := iprange.ParseList("10.0.4.1/24")

	if err != nil {
		log.Printf("error: %s", err)
	}

	log.Printf("%+v", list)

	// 然后调用 AddressRange 的 Expand方法返回一个 []net.IP
	rng := list.Expand()
	log.Printf("%s", rng)
}
