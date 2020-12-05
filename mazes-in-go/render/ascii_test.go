package render

import (
	"github.com/lithammer/dedent"
	"github.com/stretchr/testify/assert"
	"mazes/maze"
	"testing"
)

func Test_displays_simple_1x1_grid(t *testing.T) {
	g := maze.NewGrid(1, 1)

	str := ToAscii(g)

	expected := dedent.Dedent(`
		+---+
		|   |
		+---+`)
	assert.Equal(t, expected, str)
}
func Test_displays_simple_2x2_grid(t *testing.T) {
	g := maze.NewGrid(2, 2)

	str := ToAscii(g)

	expected := dedent.Dedent(`
		+---+---+
		|   |   |
		+---+---+
		|   |   |
		+---+---+`)
	assert.Equal(t, expected, str)
}

func Test_displays_complex_4x4_grid(t *testing.T) {
	g := complex4x4Grid()

	str := ToAscii(g)

	expected := dedent.Dedent(`
		+---+---+---+---+
		|               |
		+   +   +---+   +
		|   |       |   |
		+---+---+   +---+
		|               |
		+   +---+   +   +
		|       |   |   |
		+---+---+---+---+`)
	assert.Equal(t, expected, str)
}
