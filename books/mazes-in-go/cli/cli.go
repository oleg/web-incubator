package main

import (
	"flag"
	"fmt"
	"math/rand"
	"mazes"
	"time"
)

func main() {
	height := flag.Int("height", 10, "an int")
	width := flag.Int("width", 10, "an int")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	grid := mazes.NewGrid(*height, *width)
	mazes.BinaryTree(grid)
	ascii := mazes.ToAscii(grid)

	fmt.Println(ascii)
}
