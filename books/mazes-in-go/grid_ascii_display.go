package mazes

func ToAscii(grid *Grid) string {
	str := "\n" + printTop(grid, grid.cells[0]) + "\n"
	for _, row := range grid.cells {
		str += printMiddle(grid, row) + "\n"
		str += printBottom(grid, row) + "\n"
	}
	return str[:len(str)-1]
}

func printTop(grid *Grid, row []*Cell) string {
	str := "+"
	for _, cell := range row {
		str += choose(cell.linked(grid.North(cell)), "   +", "---+")
	}
	return str
}

func printBottom(grid *Grid, row []*Cell) string {
	str := "+"
	for _, cell := range row {
		str += choose(cell.linked(grid.South(cell)), "   +", "---+")
	}
	return str
}
func printMiddle(grid *Grid, row []*Cell) string {
	str := "|"
	for _, cell := range row {
		str += choose(cell.linked(grid.East(cell)), "    ", "   |")
	}
	return str
}

func choose(condition bool, t, f string) string {
	if condition {
		return t
	}
	return f
}
