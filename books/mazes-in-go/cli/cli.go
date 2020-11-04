package main

import (
	"flag"
	"fmt"
	"math/rand"
	"mazes"
	"time"
)

type Algo func(grid mazes.Grid)

var algorithms = map[string]Algo{
	"binary":     mazes.BinaryTree,
	"sidewinder": mazes.Sidewinder,
}

func main() {
	height := flag.Int("height", 10, "height of the maze")
	width := flag.Int("width", 10, "width of the maze")
	algorithm := flag.String("algorithm", "binary", "algorithm to generate maze")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	grid := mazes.NewGrid(*height, *width)
	algorithms[*algorithm](grid)
	ascii := mazes.ToAscii(grid)

	fmt.Println(ascii)
}
