package main

import (
	"fmt"
	"log"
	"net"
	"os"
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
			_, _ = fmt.Fprintf(os.Stderr, "Failed to conn connection %v\n", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to read data %v\n", err)
			return
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to write data %v\n", err)
			return
		}
	}
}
