package main

import "github.com/shirou/gopsutil/cpu"

import (
	"fmt"
	"time"
)


func main (){
//	test_CPUCounts()
	//test_CPUInof()
	test_CPUPercent()
	//test_CPUTimes()
}

func test_CPUPercent(){

	fmt.Printf("CPUPercent : ======================= start \n")
	per, err := cpu.CPUPercent(1000 * time.Millisecond, true)
	if err == nil {
		fmt.Printf("CPUPercent : %f\n", per)
	}
	fmt.Printf("CPUPercent : ======================= end \n")
}

func test_CPUCounts () {
	fmt.Printf("CPUCounts : ======================== start \n")
	counts, err := cpu.CPUCounts(false)
	if (err == nil){
		fmt.Printf("CPUCounts :", counts)
		fmt.Printf("\n")
	}
	fmt.Printf("CPUCounts : ======================== end \n")
}

func test_CPUInof() {
	fmt.Printf("CPUInfo : ======================== start \n")
	infos, err := cpu.CPUInfo()
	if (err == nil){
		for key, value := range infos {
			fmt.Printf("Key :", key)
			fmt.Printf("Value :", value)
			fmt.Printf("\n")
		}
	}
	fmt.Printf("CPUInfo : ======================== end\n")
}

func test_CPUTimes() {
	stats, err := cpu.CPUTimes(true)
	if (err == nil) {
		for key, value := range stats{
			fmt.Printf("Key :", key)
			fmt.Printf("Value :", value)
			fmt.Printf("\n")
		}
	}
}
