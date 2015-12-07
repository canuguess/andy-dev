package main

import (
	"github.com/shirou/gopsutil/net"
	"fmt"
)

func main() {
	//test_NetConnections()
	//test_NetConnectionsPid()
	test_NetIOCounters()
	//test_NetInterfaces()
}

func test_NetConnections(){
	stats, err := net.NetConnections("")

	if err != nil {
		fmt.Printf("ERROR : %s.", err)
		return
	}

	for _, stat := range stats {
		fmt.Printf("Sata : %s.\n", stat)
	}
}

func test_NetConnectionsPid(){
	stats, err := net.NetConnectionsPid("tcp", 15409)

	if err != nil {
		fmt.Printf("ERROR : %s.", err)
		return
	}

	for _, stat := range stats {
		fmt.Printf("Sata : %s.\n", stat)
	}
}

func test_NetIOCounters(){
	stats, err := net.NetIOCounters(true)

	if err != nil {
		fmt.Printf("ERROR : %s.", err)
		return
	}

	for _, stat := range stats {
		fmt.Printf("Sata : %s.\n", stat)
	}
}

func test_NetInterfaces(){
	stats, err := net.NetInterfaces()

	if err != nil {
		fmt.Printf("ERROR : %s.", err)
		return
	}

	for _, stat := range stats {
		fmt.Printf("Sata : %s.\n", stat)
	}
}
