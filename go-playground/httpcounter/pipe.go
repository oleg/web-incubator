package main

import "time"

func makePipe(logic func(time.Time) int) (chan<- time.Time, <-chan int, func()) {
	in := make(chan time.Time)
	out := make(chan int)
	pipe := func() {
		for {
			out <- logic(<-in)
		}
	}
	return in, out, pipe
}
