package main

import (
	"fmt"
	"github.com/shirou/gopsutil/load"
)

func main() {
	test_loadavg()
}

func test_loadavg() {
	stat, err := load.LoadAvg()
	if err != nil {
		fmt.Printf("ERROR : %s.\n", err)
		return
	}
	fmt.Printf("stat : %s.\n", stat.String())
}
