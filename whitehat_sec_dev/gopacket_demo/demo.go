package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

// catch local host packet and output
var (
	device         = "ens33"
	snapshotLength = 1024
	promisc        = false
	timeout        = 30 * time.Second
	handle         *pcap.Handle
	err            error
)

func main() {
	handle, err = pcap.OpenLive(device, int32(snapshotLength), promisc, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet.Dump())
	}
}
