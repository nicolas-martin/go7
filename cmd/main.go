package main

import (
	"fmt"
	"gos7"
	"log"
	"os"
	"time"
)

func main() {

	const (
		tcpDevice = "127.0.0.1"
		rack      = 0
		slot      = 2
	)
	// TCPClient
	handler := gos7.NewTCPClientHandler(tcpDevice, rack, slot)
	handler.Timeout = 200 * time.Second
	handler.IdleTimeout = 200 * time.Second
	handler.Logger = log.New(os.Stdout, "tcp: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	handler.Connect()
	defer handler.Close()
	//init client
	client := gos7.NewClient(handler)

	info, err := client.GetCPUInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)
}
