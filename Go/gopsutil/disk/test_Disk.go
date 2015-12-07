package main

import (
	"log"
	"fmt"
	"github.com/shirou/gopsutil/disk"
)

func main() {
	test_DiskIOCounters()
	//test_DiskPartitions()
	//test_DiskUsage()
}

func test_DiskIOCounters() {
	disk_io_counters, err := disk.DiskIOCounters()

	if err != nil {
		log.Printf("ERROR : %s", err)
	}

	for key, value := range disk_io_counters {
		fmt.Printf("key : %s, value : %s.\n", key, value.String())
	}
}

func test_DiskPartitions() {
	stat, err := disk.DiskPartitions(false)

	if err != nil {
		log.Printf("ERROR : %s", err)
	}

	for _, value := range stat {
		fmt.Printf("value : %s.\n", value.String())
	}
}

func test_DiskUsage() {
	stat, err := disk.DiskUsage("/home/andy")

	if err != nil {
		log.Printf("ERROR : %s", err)
	}
	fmt.Printf("value : %s.\n", stat.String())
}
