package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	name := os.Args[1]
	addr4, err := net.ResolveIPAddr("ip4", name)
	if err != nil {
		log.Fatalf("ip4 resolution error %v", err)
	}
	if addr4 == nil {
		log.Fatalf("invalid ip4 address")
	}
	fmt.Println("the ip4 address  is ", addr4.String())

	addr6, err := net.ResolveIPAddr("ip6", name)
	if err != nil {
		log.Fatalf("ip6 resolution error %v", err)
	}
	if addr6 == nil {
		log.Fatalf("invalid ip6 address")
	}
	fmt.Println("the ip6 address is ", addr6.String())
}
