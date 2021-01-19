package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp4", ":1200")
	if err != nil {
		log.Fatalf("Wrong address %v", err)
	}
	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		log.Fatalf("Failed to dial udp %v", err)
	}
	_, err = conn.Write([]byte("hello"))
	if err != nil {
		log.Fatalf("Faield to write to udp %v", err)
	}
	var buf [512]byte
	_, err = conn.Read(buf[:])
	if err != nil {
		log.Fatalf("Failed to read from udp %v", err)
	}
	fmt.Println(string(buf[:]))
}
