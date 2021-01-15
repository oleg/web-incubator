package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr := net.ParseIP(os.Args[1])
	if addr == nil {
		log.Fatalf("Invalid address")
	}
	fmt.Println("The address is ", addr.String())
}
