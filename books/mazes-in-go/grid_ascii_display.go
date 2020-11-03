package mazes

func ToAscii(grid Grid) string {
	str := "\n" + printTop(grid.cells[0]) + "\n"
	for _, row := range grid.cells {
		str += printMiddle(row) + "\n"
		str += printBottom(row) + "\n"
	}
	return str[:len(str)-1]
}

func printTop(row []*Cell) string {
	str := "+"
	for _, cell := range row {
		str += choose(cell.linked(cell.north), "   +", "---+")
	}
	return str
}

func printBottom(row []*Cell) string {
	str := "+"
	for _, cell := range row {
		str += choose(cell.linked(cell.south), "   +", "---+")
	}
	return str
}
func printMiddle(row []*Cell) string {
	str := "|"
	for _, cell := range row {
		str += choose(cell.linked(cell.east), "    ", "   |")
	}
	return str
}

func choose(condition bool, t, f string) string {
	if condition {
		return t
	}
	return f
}
