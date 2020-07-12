package mazes_in_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_creates_new_grid_with_100_cells(t *testing.T) {
	grid := NewGrid(10, 10)
	count := 0
	for _, row := range grid.cells {
		for range row {
			count++
		}
	}
	assert.Equal(t, 100, count)
}
