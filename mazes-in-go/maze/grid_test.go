package maze

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sets_correct_rows_and_columns(t *testing.T) {
	grid := NewGrid(2, 3)

	assert.Equal(t, 0, grid.cells[0][0].row)
	assert.Equal(t, 0, grid.cells[0][0].column)

	assert.Equal(t, 0, grid.cells[0][1].row)
	assert.Equal(t, 1, grid.cells[0][1].column)

	assert.Equal(t, 0, grid.cells[0][2].row)
	assert.Equal(t, 2, grid.cells[0][2].column)

	assert.Equal(t, 1, grid.cells[1][0].row)
	assert.Equal(t, 0, grid.cells[1][0].column)

	assert.Equal(t, 1, grid.cells[1][1].row)
	assert.Equal(t, 1, grid.cells[1][1].column)

	assert.Equal(t, 1, grid.cells[1][2].row)
	assert.Equal(t, 2, grid.cells[1][2].column)
}

func Test_sets_north_south_east_west_sides(t *testing.T) {
	grid := NewGrid(3, 3)

	a := grid.cells[0][0]
	b := grid.cells[0][1]
	c := grid.cells[0][2]

	d := grid.cells[1][0]
	e := grid.cells[1][1]
	f := grid.cells[1][2]

	g := grid.cells[2][0]
	h := grid.cells[2][1]
	i := grid.cells[2][2]

	n := (*Cell)(nil)

	//	a|b|c
	//	d|e|f
	//	g|h|i

	assertNorthEastSouthWest(t, a, n, b, d, n, grid)
	assertNorthEastSouthWest(t, b, n, c, e, a, grid)
	assertNorthEastSouthWest(t, c, n, n, f, b, grid)

	assertNorthEastSouthWest(t, d, a, e, g, n, grid)
	assertNorthEastSouthWest(t, e, b, f, h, d, grid)
	assertNorthEastSouthWest(t, f, c, n, i, e, grid)

	assertNorthEastSouthWest(t, g, d, h, n, n, grid)
	assertNorthEastSouthWest(t, h, e, i, n, g, grid)
	assertNorthEastSouthWest(t, i, f, n, n, h, grid)
}

func assertNorthEastSouthWest(t *testing.T, c, north, east, south, west *Cell, g *Grid) {
	assert.Equal(t, north, g.North(c))
	assert.Equal(t, east, g.East(c))
	assert.Equal(t, south, g.South(c))
	assert.Equal(t, west, g.West(c))
}

func Test_creates_new_grid_with_100_cells(t *testing.T) {
	grid := NewGrid(3, 7)

	assert.Len(t, grid.cells, 3)
	for _, row := range grid.cells {
		assert.Len(t, row, 7)
	}
}

func Test_iterate_over_all_cells_correct_count(t *testing.T) {
	grid := NewGrid(7, 7)

	count := 0
	grid.EachCells(func(cell *Cell) {
		count++
	})

	assert.Equal(t, 49, count)
}

func Test_access_returns_cell_for_correct_coordinates(t *testing.T) {
	grid := NewGrid(4, 7)

	for i, row := range grid.cells {
		for j, cell := range row {
			assert.Equal(t, cell, grid.Cell(i, j))
		}
	}
}

func Test_access_returns_nil_for_incorrect_coordinates(t *testing.T) {
	grid := NewGrid(4, 7)

	assert.Nil(t, grid.Cell(-1, 0))
	assert.Nil(t, grid.Cell(0, -1))
	assert.Nil(t, grid.Cell(4, 0))
	assert.Nil(t, grid.Cell(0, 7))
	assert.Nil(t, grid.Cell(5, 0))
	assert.Nil(t, grid.Cell(0, 8))
}

func Test_iterate_over_all_cells_correct_cells(t *testing.T) {
	grid := NewGrid(3, 2)
	visitedCells := make([]*Cell, 0)

	grid.EachCells(func(cell *Cell) {
		visitedCells = append(visitedCells, cell)
	})

	assert.Len(t, visitedCells, 6)

	assert.Equal(t, grid.cells[0][0], visitedCells[0])
	assert.Equal(t, grid.cells[0][1], visitedCells[1])

	assert.Equal(t, grid.cells[1][0], visitedCells[2])
	assert.Equal(t, grid.cells[1][1], visitedCells[3])

	assert.Equal(t, grid.cells[2][0], visitedCells[4])
	assert.Equal(t, grid.cells[2][1], visitedCells[5])
}

func Test_grid_returns_north_cell(t *testing.T) {
	grid := NewGrid(4, 4)

	assert.Nil(t, grid.North(grid.Cell(0, 0)))
	assert.Nil(t, grid.North(grid.Cell(0, 1)))
	assert.Nil(t, grid.North(grid.Cell(0, 2)))
	assert.Nil(t, grid.North(grid.Cell(0, 3)))

	AssertEqualCell(t, grid.Cell(0, 0), grid.North(grid.Cell(1, 0)))
	AssertEqualCell(t, grid.Cell(0, 1), grid.North(grid.Cell(1, 1)))
	AssertEqualCell(t, grid.Cell(0, 2), grid.North(grid.Cell(1, 2)))
	AssertEqualCell(t, grid.Cell(0, 3), grid.North(grid.Cell(1, 3)))

	AssertEqualCell(t, grid.Cell(1, 0), grid.North(grid.Cell(2, 0)))
	AssertEqualCell(t, grid.Cell(1, 1), grid.North(grid.Cell(2, 1)))
	AssertEqualCell(t, grid.Cell(1, 2), grid.North(grid.Cell(2, 2)))
	AssertEqualCell(t, grid.Cell(1, 3), grid.North(grid.Cell(2, 3)))

	AssertEqualCell(t, grid.Cell(2, 0), grid.North(grid.Cell(3, 0)))
	AssertEqualCell(t, grid.Cell(2, 1), grid.North(grid.Cell(3, 1)))
	AssertEqualCell(t, grid.Cell(2, 2), grid.North(grid.Cell(3, 2)))
	AssertEqualCell(t, grid.Cell(2, 3), grid.North(grid.Cell(3, 3)))
}

func Test_grid_returns_east_cell(t *testing.T) {
	grid := NewGrid(4, 4)

	AssertEqualCell(t, grid.Cell(0, 1), grid.East(grid.Cell(0, 0)))
	AssertEqualCell(t, grid.Cell(1, 1), grid.East(grid.Cell(1, 0)))
	AssertEqualCell(t, grid.Cell(2, 1), grid.East(grid.Cell(2, 0)))
	AssertEqualCell(t, grid.Cell(3, 1), grid.East(grid.Cell(3, 0)))

	AssertEqualCell(t, grid.Cell(0, 2), grid.East(grid.Cell(0, 1)))
	AssertEqualCell(t, grid.Cell(1, 2), grid.East(grid.Cell(1, 1)))
	AssertEqualCell(t, grid.Cell(2, 2), grid.East(grid.Cell(2, 1)))
	AssertEqualCell(t, grid.Cell(3, 2), grid.East(grid.Cell(3, 1)))

	AssertEqualCell(t, grid.Cell(0, 3), grid.East(grid.Cell(0, 2)))
	AssertEqualCell(t, grid.Cell(1, 3), grid.East(grid.Cell(1, 2)))
	AssertEqualCell(t, grid.Cell(2, 3), grid.East(grid.Cell(2, 2)))
	AssertEqualCell(t, grid.Cell(3, 3), grid.East(grid.Cell(3, 2)))

	assert.Nil(t, grid.East(grid.Cell(0, 3)))
	assert.Nil(t, grid.East(grid.Cell(1, 3)))
	assert.Nil(t, grid.East(grid.Cell(2, 3)))
	assert.Nil(t, grid.East(grid.Cell(3, 3)))
}

func Test_grid_returns_south_cell(t *testing.T) {
	grid := NewGrid(4, 4)

	AssertEqualCell(t, grid.Cell(1, 0), grid.South(grid.Cell(0, 0)))
	AssertEqualCell(t, grid.Cell(1, 1), grid.South(grid.Cell(0, 1)))
	AssertEqualCell(t, grid.Cell(1, 2), grid.South(grid.Cell(0, 2)))
	AssertEqualCell(t, grid.Cell(1, 3), grid.South(grid.Cell(0, 3)))

	AssertEqualCell(t, grid.Cell(2, 0), grid.South(grid.Cell(1, 0)))
	AssertEqualCell(t, grid.Cell(2, 1), grid.South(grid.Cell(1, 1)))
	AssertEqualCell(t, grid.Cell(2, 2), grid.South(grid.Cell(1, 2)))
	AssertEqualCell(t, grid.Cell(2, 3), grid.South(grid.Cell(1, 3)))

	AssertEqualCell(t, grid.Cell(3, 0), grid.South(grid.Cell(2, 0)))
	AssertEqualCell(t, grid.Cell(3, 1), grid.South(grid.Cell(2, 1)))
	AssertEqualCell(t, grid.Cell(3, 2), grid.South(grid.Cell(2, 2)))
	AssertEqualCell(t, grid.Cell(3, 3), grid.South(grid.Cell(2, 3)))

	assert.Nil(t, grid.South(grid.Cell(3, 0)))
	assert.Nil(t, grid.South(grid.Cell(3, 1)))
	assert.Nil(t, grid.South(grid.Cell(3, 2)))
	assert.Nil(t, grid.South(grid.Cell(3, 3)))
}

func Test_grid_returns_west_cell(t *testing.T) {
	grid := NewGrid(4, 4)

	assert.Nil(t, grid.West(grid.Cell(0, 0)))
	assert.Nil(t, grid.West(grid.Cell(1, 0)))
	assert.Nil(t, grid.West(grid.Cell(2, 0)))
	assert.Nil(t, grid.West(grid.Cell(3, 0)))

	AssertEqualCell(t, grid.Cell(0, 0), grid.West(grid.Cell(0, 1)))
	AssertEqualCell(t, grid.Cell(1, 0), grid.West(grid.Cell(1, 1)))
	AssertEqualCell(t, grid.Cell(2, 0), grid.West(grid.Cell(2, 1)))
	AssertEqualCell(t, grid.Cell(3, 0), grid.West(grid.Cell(3, 1)))

	AssertEqualCell(t, grid.Cell(0, 1), grid.West(grid.Cell(0, 2)))
	AssertEqualCell(t, grid.Cell(1, 1), grid.West(grid.Cell(1, 2)))
	AssertEqualCell(t, grid.Cell(2, 1), grid.West(grid.Cell(2, 2)))
	AssertEqualCell(t, grid.Cell(3, 1), grid.West(grid.Cell(3, 2)))

	AssertEqualCell(t, grid.Cell(0, 2), grid.West(grid.Cell(0, 3)))
	AssertEqualCell(t, grid.Cell(1, 2), grid.West(grid.Cell(1, 3)))
	AssertEqualCell(t, grid.Cell(2, 2), grid.West(grid.Cell(2, 3)))
	AssertEqualCell(t, grid.Cell(3, 2), grid.West(grid.Cell(3, 3)))
}

func AssertEqualCell(t *testing.T, expected, actual *Cell) {
	assert.Equal(t, expected.row, actual.row)
	assert.Equal(t, expected.column, actual.column)
}
