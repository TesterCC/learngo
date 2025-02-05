package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)
// local test: curl http://192.168.80.1:3001/test
// 设置数据包捕获所需要的几个变量。
var (
	iface    = "ens36"    // 接口名称
	snaplen  = int32(1600) // 快照长度：每个帧要捕获的数据量
	promisc  = false       // 决定是否以混杂模式运行
	timeout  = pcap.BlockForever   // 超时设置
	filter   = "tcp and port 3001" // BPF过滤器
	devFound = false
)
// 新找网卡，再抓包
func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}
    // 判断设备是否存在
	for _, device := range devices {
		if device.Name == iface {
			devFound = true
		}
	}
	if !devFound {
		log.Panicf("Device named '%s' does not exist\n", iface)
	}
	// capture logic, 首先要获取或创建一个*pcap.Handle，它允许我们读取和注入数据包。使用该handle，可以应用一个BPF过滤器并创建一个新的包数据源，从中读取数据包。
	handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
	if err != nil {
		log.Panicln(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(filter); err != nil {
		log.Panicln(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		fmt.Println(packet)
	}
}
