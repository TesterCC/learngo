package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

/* ref network_monitor.py
gopsutil

doc: https://pkg.go.dev/github.com/shirou/gopsutil#section-documentation

ref:
1. https://www.cnblogs.com/nickchen121/p/11517451.html
2. https://blog.csdn.net/weixin_43842373/article/details/115632875

install：
go get github.com/shirou/gopsutil

for devops and suspicious process analysis
*/

// 1 - CPU
// get cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("Get CPU info failed, err : %v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}

	// CPU usage rate
	// for{} 实现无限循环， 相当于python的while True
	for {
		percent, _ := cpu.Percent(time.Second, true)
		fmt.Printf("CPU percent: %v \n", percent)
	}
}

// get cpu load info
func getCpuLoad() {
	info, _ := load.Avg()
	fmt.Printf("CPU Load info: %v\n", info) // 360需要添加信任
}

// 2 - Memory
// get memory info
func getMemoryInfo() {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("Memory info : %v", memInfo)
}

// 3 - Host
// maybe need install package:
// go get -u golang.org/x/sys/windows
func getHostInfo() {
	hi, _ := host.Info()
	fmt.Printf("host info: %v\n update time: %v\n boot time: %v\n", hi, hi.Uptime, hi.BootTime)
	fmt.Println(hi.Hostname)
	fmt.Println(hi.HostID)
}

// 4 - Disk

func getDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Println("get Partitions failed, err: %v\n", err)
		return
	}

	// go for loop
	for _, part := range parts {
		fmt.Printf("part: %v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info: used: %v\nfree: %v\n", diskInfo.UsedPercent, diskInfo.Free)
	}

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}

}

// 5 - network IO

func getNetInfo()  {
	info,_ := net.IOCounters(true)
	for index,v := range info{
		fmt.Printf("%v:%v send:%v rece:%v\n", index,v,v.BytesSent,v.BytesRecv)
	}
}


// 6 - Process  参考文档

func main() {
	//getCpuInfo()
	//fmt.Println("--------------")
	//getCpuLoad()
	//getMemoryInfo()
	//getHostInfo()
	//getDiskInfo()
	getNetInfo()
}
