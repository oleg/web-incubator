package mazes

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

	assertNorthEastSouthWest(t, a, n, b, d, n)
	assertNorthEastSouthWest(t, b, n, c, e, a)
	assertNorthEastSouthWest(t, c, n, n, f, b)

	assertNorthEastSouthWest(t, d, a, e, g, n)
	assertNorthEastSouthWest(t, e, b, f, h, d)
	assertNorthEastSouthWest(t, f, c, n, i, e)

	assertNorthEastSouthWest(t, g, d, h, n, n)
	assertNorthEastSouthWest(t, h, e, i, n, g)
	assertNorthEastSouthWest(t, i, f, n, n, h)
}

func assertNorthEastSouthWest(t *testing.T, c, north, east, south, west *Cell) {
	assert.Equal(t, north, c.north)
	assert.Equal(t, east, c.east)
	assert.Equal(t, south, c.south)
	assert.Equal(t, west, c.west)
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
