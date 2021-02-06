package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	timestamps := Timestamps{
		Interval:         10 * time.Second, //todo make configurable? via command line?
		CleanupThreshold: 10_000,           //depends on load nature
	}
	in, out, pipe := makePipe(timestamps.AddAndCount)
	go pipe()

	log.Fatal(http.ListenAndServe(":8080", &CounterHandler{
		timestamps: in,
		counts:     out,
	}))
}
