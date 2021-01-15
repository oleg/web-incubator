package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {

	address := os.Args[1]
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Fatalf("Wrong address %v", err)
	}
	tcp, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalf("Failed to dial tcp %v", err)
	}
	_, err = tcp.Write([]byte("HEAD / HTTP/1.0 \r\n\r\n"))
	if err != nil {
		log.Fatalf("Faield to write to tcp %v", err)
	}
	response, err := ioutil.ReadAll(tcp)
	if err != nil {
		log.Fatalf("Failed to read from tcp %v", err)
	}
	fmt.Println(string(response))
}
