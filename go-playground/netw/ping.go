package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	laddr, err := net.ResolveIPAddr("ip4", "0.0.0.0")
	if err != nil {
		fmt.Println("Local address resolution error", err)
		os.Exit(1)
	}
	raddr, err := net.ResolveIPAddr("ip4", os.Args[1])
	if err != nil {
		fmt.Println("Remote address resolution error", err)
		os.Exit(1)
	}
	conn, err := net.DialIP("ip4:icmp", laddr, raddr)
	if err != nil {
		fmt.Println("Dial error", err)
		os.Exit(1)
	}
	msg := [512]byte{
		0: 8,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 13,
		6: 0,
		7: 37,
	}
	l := 8
	check := checkSum(msg[:l])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	_, err = conn.Write(msg[:l])
	if err != nil {
		fmt.Println("Write error", err)
		os.Exit(1)
	}

	fmt.Println("Message sent:")
	for n := 0; n < 8; n++ {
		fmt.Print(" ", msg[n])
	}
	fmt.Println()

	size, err := conn.Read(msg[:])
	if err != nil {
		fmt.Println("Read error", err)
		os.Exit(1)
	}
	fmt.Println("Message received:")
	ipv4HeaderSize := 20
	for n := ipv4HeaderSize; n < size; n++ {
		fmt.Print(" ", msg[n])
	}
	fmt.Println()
}

func checkSum(msg []byte) uint16 {
	sum := 0
	for n := 0; n < len(msg); n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += sum >> 16
	return uint16(^sum)
}
