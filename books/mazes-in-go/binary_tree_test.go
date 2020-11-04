package mazes

import (
	"github.com/lithammer/dedent"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_produces_binary_tree(t *testing.T) {
	rand.Seed(42)
	grid := NewGrid(5, 5)

	BinaryTree(grid)

	expected := dedent.Dedent(`
	+---+---+---+---+---+
	|                   |
	+---+---+   +   +   +
	|           |   |   |
	+---+---+---+   +   +
	|               |   |
	+   +---+---+---+   +
	|   |               |
	+---+   +   +---+   +
	|       |   |       |
	+---+---+---+---+---+`)
	assert.Equal(t, expected, ToAscii(grid))
}

func Test_returns_nil_if_north_and_east_are_nil(t *testing.T) {
	grid := NewGrid(5, 5)
	cell := grid.Cell(0, 4)

	neighbor := chooseNeighbor(cell)

	assert.Nil(t, neighbor)
}

func Test_returns_north_if_east_is_nil(t *testing.T) {
	grid := NewGrid(5, 5)
	cell := grid.Cell(1, 4)

	neighbor := chooseNeighbor(cell)

	assert.Equal(t, grid.Cell(0, 4), neighbor)
}

func Test_returns_east_if_north_is_nil(t *testing.T) {
	grid := NewGrid(5, 5)
	cell := grid.Cell(0, 3)

	neighbor := chooseNeighbor(cell)

	assert.Equal(t, grid.Cell(0, 4), neighbor)
}

//todo hangs on compare assert.Equal(t, grid.Cell(3, 3), grid.Cell(2, 3))

func Test_returns_rand_if_east_and_north_are_not_nil(t *testing.T) {
	grid := NewGrid(5, 5)
	cell := grid.Cell(3, 3)

	rand.Seed(42)
	neighbor := chooseNeighbor(cell)

	assert.Equal(t, grid.Cell(3, 4), neighbor)
}
