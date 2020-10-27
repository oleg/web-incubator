package mazes

type Grid struct {
	height, width int
	cells         [][]*Cell //todo one dimensional array?
}

func NewGrid(height, width int) Grid {
	grid := Grid{height, width, makeCells(height, width)}
	initCells(&grid)
	return grid
}

func makeCells(height int, width int) [][]*Cell {
	cells := make([][]*Cell, height)
	for i := range cells {
		row := make([]*Cell, width)
		for j := 0; j < width; j++ {
			row[j] = NewCell(i, j)
		}
		cells[i] = row
	}
	return cells
}

func initCells(g *Grid) {
	for i, row := range g.cells {
		for j, cell := range row {
			cell.north = g.Cell(i-1, j)
			cell.east = g.Cell(i, j+1)
			cell.south = g.Cell(i+1, j)
			cell.west = g.Cell(i, j-1)
		}
	}
}

func (g *Grid) EachCells(f func(cell *Cell)) {
	for _, row := range g.cells {
		for _, cell := range row {
			f(cell)
		}
	}
}

func (g *Grid) Cell(row int, column int) *Cell {
	if row < 0 || row >= g.height {
		return nil
	}
	if column < 0 || column >= g.width {
		return nil
	}
	return g.cells[row][column]
}
