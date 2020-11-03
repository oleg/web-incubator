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
			+---+---+---+   +   +
			|               |   |
			+---+---+---+---+   +
			|                   |
			+   +---+   +---+   +
			|   |       |       |
			+   +   +   +   +   +
			|   |   |   |   |   |
			+---+---+---+---+---+`)
	assert.Equal(t, expected, ToAscii(grid))
}
