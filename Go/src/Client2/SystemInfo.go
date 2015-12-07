package main

import (
	"log"
	"time"
	"os"
	"strings"
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/net"
)

var
(
	ServerId = 1
)

func getSystemInfo() (SystemInfo) {
	info := SystemInfo{}
	info.collectInfo()
	return info
}

type SystemInfo struct {
	ServerId	int		`json:server_id`
	PerCPU		float64 `json:percpu`
	Memoey		uint64	`json:memory`
	PerMemory	float64	`json:permemory`
	DskRead		uint64	`json:dsk_read`
	DskWrite	uint64	`json:dsk_write`
	NetRecv		uint64	`json:net_recv`
	NetSend		uint64	`json:net_send`
}

func (info *SystemInfo) collectInfo() {
	logger.Printf("<%d:%s>")
	info.ServerId = ServerId

	info.collectMemInfo()
	info.collectDiskInfo()
	info.collectNetInfo()

	info.collectCPUInfo()

	info.collectDiskInfo()
	info.collectNetInfo()
}

func (info *SystemInfo) collectCPUInfo() {
	per, err := cpu.CPUPercent(1000 * time.Millisecond, true)
	checkError(err)
	info.PerCPU = per[0]
}

func (info *SystemInfo) collectMemInfo() {
	stat, err := mem.SwapMemory()
	checkError(err)
	info.Memoey = stat.Used;
	info.PerMemory = float64(info.Memoey) /float64(stat.Total);
}

func (info *SystemInfo) collectDiskInfo() {
	if info.DskRead == 0 {
		disk_io_counters, err := disk.DiskIOCounters()
		checkError(err)
		for key, value := range disk_io_counters {
			if len(key) == 3 {
				info.DskWrite += value.WriteBytes
				info.DskRead += value.ReadBytes
			}
		}
	} else {
		write_bytes := info.DskWrite
		read_bytes := info.DskRead
		disk_io_counters, err := disk.DiskIOCounters()
		checkError(err)
		for key, value := range disk_io_counters {
			if len(key) == 3 {
				info.DskWrite += value.WriteBytes
				info.DskRead += value.ReadBytes
			}
		}
		info.DskWrite -= write_bytes
		info.DskRead -= read_bytes
	}
}

func (info *SystemInfo) collectNetInfo() {
	if info.DskRead == 0 {
		stats, err := net.NetIOCounters(true)
		checkError(err)

		for _, value := range stats {
			if strings.HasPrefix(value.Name, "eth") {
				info.NetSend += value.BytesSent
				info.NetRecv += value.BytesRecv
			}
		}
	} else {
		sent_bytes := info.NetSend
		recv_bytes := info.NetRecv
		stats, err := net.NetIOCounters(true)
		checkError(err)

		for _, value := range stats {
			if strings.HasPrefix(value.Name, "eth") {
				info.NetSend += value.BytesSent
				info.NetRecv += value.BytesRecv
			}
		}
		info.NetSend -= sent_bytes
		info.NetRecv -= recv_bytes
	}
}

func (info SystemInfo) String() string {
	s, _ := json.Marshal(info)
	return string(s)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("[ERROR] : %s", err)
		os.Exit(-1)
	}
}

