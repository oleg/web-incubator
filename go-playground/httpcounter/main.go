package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const storageFilename = ".timestamps"
const defaultDuration = 1 * time.Minute

func main() {
	repo, err := NewFileTimestampsRepo(storageFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()

	counter, err := NewTimestampCounter(countWindow(), repo)
	if err != nil {
		log.Fatal(err)
	}

	in, out, pipe := makePipe(counter.AddAndCount)
	go pipe()

	err = http.ListenAndServe(":"+port(), &CounterHandler{
		timestamps: in,
		counts:     out,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func countWindow() time.Duration {
	countWindow := os.Getenv("TS_COUNT_WINDOW_SEC")
	dur, err := strconv.Atoi(countWindow)
	if err == nil {
		return time.Duration(dur) * time.Second
	}
	return defaultDuration
}

func port() string {
	port := os.Getenv("TS_PORT")
	if port != "" {
		return port
	}
	return "8080"
}
