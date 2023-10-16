package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

//P176-177  8.2 使用pcap子包识别设备
func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}

	for _, device := range devices {
		fmt.Printf("Device Name: %s\n", device.Name)   // ens33, ens36 is ""
		fmt.Printf("Desc: %s\n", device.Description)
		fmt.Printf("Flags: %d\n", device.Flags)
		fmt.Println("------------------------------------------")

		// get ip info by device.Addresses
		for _, address := range device.Addresses {
			fmt.Printf("    IP:      %s\n", address.IP)
			fmt.Printf("    Netmask: %s\n", address.Netmask)
		}
		fmt.Println("===========================================")
	}
}
