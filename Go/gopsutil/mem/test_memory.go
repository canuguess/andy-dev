package main

import (
	"github.com/shirou/gopsutil/mem"
	"fmt"
)

func main() {
	test_SwapMemoryStat()
	test_VirtualMemoryStat()
}

func test_SwapMemoryStat() {
	stat, err := mem.SwapMemory()
	if err == nil {
		fmt.Printf(stat.String())
		fmt.Printf("\n")
	}
}

func test_VirtualMemoryStat() {
	stat, err := mem.VirtualMemory()
	if err == nil {
		fmt.Printf(stat.String())
		fmt.Printf("\n")
	}
}
