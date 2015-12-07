package main

import (
	"flag"
	"log"
	"net"
	"os"
)

var (
	logger = log.New(os.Stdout, "[UDP-SERVER]", log.LstdFlags)
	host = flag.String("host", "0.0.0.0", "host to listen on")
	port = flag.Int("port", 18888, "port to listen on")
	block_size = flag.Int("size", 1024, "block size to read packets on")
)


func main(){
	flag.Parse()

	ip := net.ParseIP(*host)
	listener, err := net.ListenUDP("udp", &net.UDPAddr{ IP : ip, Port : *port})
	if err != nil {
		logger.Fatalf("error during listen : %s", err)
	}
	logger.Printf("listening on addr=%s with block size=%d", listener.LocalAddr(), *block_size)

	buffer := make([]byte, *block_size)

	for {
		n, remoteAddr, err := listener.ReadFromUDP(buffer)
		if err != nil {
			logger.Fatalf("error during read : %s", err)
		}
		logger.Printf("<%s>%d:%s", remoteAddr, n, buffer[0:n])
	}
}

