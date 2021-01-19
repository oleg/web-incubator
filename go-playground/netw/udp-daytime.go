package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp4", ":1200")
	if err != nil {
		log.Fatalf("Failed to resolve addr: %v", err)
	}
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatalf("Failed to listen addr %v: %v", addr, err)
	}
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[:])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to read from UDP %v\n", err)
		return
	}

	daytime := time.Now().String()
	_, err = conn.WriteTo([]byte(daytime), addr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to write to UDP %v\n", err)
		return
	}
}
