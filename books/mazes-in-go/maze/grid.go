package maze

type Grid struct {
	height, width int
	cells         [][]*Cell
}

func NewGrid(height, width int) *Grid {
	return &Grid{height, width, makeCells(height, width)}
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

////todo:oleg test
//func (g *Grid) First() *Cell {
//	return g.cells[0]
//}

func (g *Grid) EachCells(f func(cell *Cell)) {
	for _, row := range g.cells {
		for _, cell := range row {
			f(cell)
		}
	}
}

//todo:oleg test me
func (g *Grid) EachRow(f func(n int, row []*Cell)) {
	for n, row := range g.cells {
		f(n, row)
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

func (g *Grid) North(cell *Cell) *Cell { //*Cell, bool?
	return g.Cell(cell.row-1, cell.column)
}

func (g *Grid) East(cell *Cell) *Cell { //*Cell, bool?
	return g.Cell(cell.row, cell.column+1)
}

func (g *Grid) South(cell *Cell) *Cell { //*Cell, bool?
	return g.Cell(cell.row+1, cell.column)
}

func (g *Grid) West(cell *Cell) *Cell { //*Cell, bool?
	return g.Cell(cell.row, cell.column-1)
}
