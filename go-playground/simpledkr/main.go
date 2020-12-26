package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "Hello %s!", path.Base(request.URL.Path))
	if err != nil {
		log.Printf("Error %v\n", err)
	}
}
