package main

import (
	"flag"
	"log"
	"net"
	"os"
	"time"
	"bytes"
)

var (
	logger = log.New(os.Stdout, "[UDP-Client]", log.LstdFlags)
	host = flag.String("host", "192.168.99.184", "host to send udp to")
	port = flag.Int("port", 18888, "port to send udp to")
)

func main() {
	flag.Parse()

	ip := net.ParseIP(*host)

	srcAddr := &net.UDPAddr{IP : []byte{0,0,0,0}, Port : 0}
	dstAddr := &net.UDPAddr{IP : ip, Port : *port}

	logger.Printf("sending from %s to %s", srcAddr, dstAddr)

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		logger.Fatal(err)
	}

	for {
		info := getSystemInfo()
		buffer := bytes.NewBufferString(info.String());
	//	n, err := conn.Write(buffer.Bytes())
	//	if err != nil {
	//		logger.Fatal(err)
	//	}
		time.Sleep(200 * time.Millisecond)
		logger.Printf("<%d:%s>", 10, buffer)
	}
	defer conn.Close()
}

