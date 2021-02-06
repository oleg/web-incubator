package main

import (
	"fmt"
	"net/http"
	"time"
)

type CounterHandler struct {
	timestamps chan<- time.Time
	counts     <-chan int
}

func (c *CounterHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	c.timestamps <- time.Now()
	count := <-c.counts
	fmt.Fprintf(w, "%d", count)
}
