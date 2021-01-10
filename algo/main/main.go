package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)
	intervals := make([]interval, n)
	for i := 0; i < n; i++ {
		var from, to int
		fmt.Fscan(os.Stdin, &from, &to)
		intervals[i] = interval{from: from, to: to}
	}
}

type interval struct {
	from, to int
}
