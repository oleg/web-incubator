package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp4", ":1200")
	if err != nil {
		log.Fatalf("Failed to resolve addr: %v", err)
	}
	tcp, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		log.Fatalf("Failed to listen addr %v: %v", addr, err)
	}
	for {
		conn, err := tcp.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to conn connection %v\n", err)
			continue
		}
		conn.Write([]byte(time.Now().String()))
		conn.Close()
	}
}
