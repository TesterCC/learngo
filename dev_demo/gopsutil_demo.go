package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"time"
)

/* ref network_monitor.py
gopsutil

doc: https://pkg.go.dev/github.com/shirou/gopsutil#section-documentation

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
func getMemoryInfo()  {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("Memory info : %v", memInfo)
}

// 3 - Host


// 4 - Disk


// 5 - network IO


// 6 - Process  参考文档

func main() {
	//getCpuInfo()
	//fmt.Println("--------------")
	//getCpuLoad()
	getMemoryInfo()
}
