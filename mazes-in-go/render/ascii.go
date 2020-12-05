package render

import (
	"mazes/maze"
)

func ToAscii(grid *maze.Grid) string {
	str := ""
	grid.EachRow(func(n int, row []*maze.Cell) {
		if n == 0 {
			str = "\n" + printTop(grid, row) + "\n"
		}
		str += printMiddle(grid, row) + "\n"
		str += printBottom(grid, row) + "\n"
	})
	return str[:len(str)-1]
}

func printTop(grid *maze.Grid, row []*maze.Cell) string {
	str := "+"
	for _, cell := range row {
		str += choose(cell.Linked(grid.North(cell)), "   +", "---+")
	}
	return str
}

func printBottom(grid *maze.Grid, row []*maze.Cell) string {
	str := "+"
	for _, cell := range row {
		str += choose(cell.Linked(grid.South(cell)), "   +", "---+")
	}
	return str
}
func printMiddle(grid *maze.Grid, row []*maze.Cell) string {
	str := "|"
	for _, cell := range row {
		str += choose(cell.Linked(grid.East(cell)), "    ", "   |")
	}
	return str
}

func choose(condition bool, t, f string) string {
	if condition {
		return t
	}
	return f
}
