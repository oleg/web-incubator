package mazes_in_go

type Grid struct {
	//or public?
	width, height int
	cells         [][]Cell
}

func NewGrid(width, height int) Grid {
	cells := make([][]Cell, width)
	for i := range cells {
		cells[i] = make([]Cell, height)
	}
	return Grid{width, height, cells}
}
