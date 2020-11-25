package main

import (
	"flag"
	"fmt"
	"math/rand"
	"mazes/generator"
	"mazes/maze"
	"mazes/render"
	"time"
)

type Algo func(grid *maze.Grid)

var algorithms = map[string]Algo{
	"binary":     generator.BinaryTree,
	"sidewinder": generator.Sidewinder,
}

func main() {
	height := flag.Int("height", 10, "height of the maze")
	width := flag.Int("width", 10, "width of the maze")
	algorithm := flag.String("algorithm", "binary", "algorithm to generate maze")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	grid := maze.NewGrid(*height, *width)
	algorithms[*algorithm](grid)
	ascii := render.ToAscii(grid)

	fmt.Println(ascii)
}
