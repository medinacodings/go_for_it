package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

func Diagnosis() {
	// Gathering CPU info
	cpuInfo, _ := cpu.Percent(time.Second, false)
	fmt.Println("CPU Usage: %.2f%%\n", cpuInfo[0])

	// Gathering Memory info
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("Total Memory: %v MB\n", memInfo.Total/1024/1024)
	fmt.Printf("Used Memory: %v MB\n", memInfo.Used/1024/1024)
}

func main() {
	// Let's check
	Diagnosis()
}
