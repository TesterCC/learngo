package main

import (
	"github.com/malfunkt/iprange"
	"log"
)

// 白帽子安全开放实战 P17-18

// 实现对多IP的扫描，引入第三方包 github.com/malfunkt/iprange   安装：go get github.com/malfunkt/iprange
// iprange实现了对多个IP的解析

// iprange 示例
func main() {
	//list, err := iprange.ParseList("10.0.0.1, 10.0.0.5-10, 192.168.1.*,192.168.10.0/24")
	//list, err := iprange.ParseList("10.0.0.1, 10.0.0.5-10")

	//list, err := iprange.ParseList("192.168.1.*,192.168.10.0/24")

	list, err := iprange.ParseList("10.0.4.1/24")

	if err != nil {
		log.Printf("error: %s", err)
	}

	log.Printf("%+v", list)

	rng := list.Expand()
	log.Printf("%s", rng)
}
