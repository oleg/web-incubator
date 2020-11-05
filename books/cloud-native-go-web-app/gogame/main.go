package main

import (
	"github.com/oleg/incubator/books/cloud-native-go-web-app/gogame/service"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "3000"
	}

	server := service.NewServer()
	server.Run(":" + port)
}
