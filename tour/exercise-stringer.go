package main

import "fmt"

/*

http://127.0.0.1:3999/methods/18
https://tour.go-zh.org/methods/18

tour solution:
https://github.com/golang/tour/blob/master/solutions/stringers.go

*/

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d",ip[0],ip[1],ip[2],ip[3])  // return string
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
		"test": {11,11,11,11},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
