package mazes

import (
	"github.com/lithammer/dedent"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_displays_simple_1x1_grid(t *testing.T) {
	g := NewGrid(1, 1)

	str := asAscii(g)

	expected := dedent.Dedent(`
		+---+
		|   |
		+---+`)
	assert.Equal(t, expected, str)
}
func Test_displays_simple_2x2_grid(t *testing.T) {
	g := NewGrid(2, 2)

	str := asAscii(g)

	expected := dedent.Dedent(`
		+---+---+
		|   |   |
		+---+---+
		|   |   |
		+---+---+`)
	assert.Equal(t, expected, str)
}

func Test_displays_complex_4x4_grid(t *testing.T) {
	g := NewGrid(4, 4)
	//0>
	g.Cell(0, 0).link(g.Cell(0, 1))
	g.Cell(0, 1).link(g.Cell(0, 2))
	g.Cell(0, 2).link(g.Cell(0, 3))
	//0V
	g.Cell(0, 0).link(g.Cell(1, 0))
	g.Cell(0, 1).link(g.Cell(1, 1))
	g.Cell(0, 3).link(g.Cell(1, 3))
	//1>
	g.Cell(1, 1).link(g.Cell(1, 2))
	//1V
	g.Cell(1, 2).link(g.Cell(2, 2))
	//2>
	g.Cell(2, 0).link(g.Cell(2, 1))
	g.Cell(2, 1).link(g.Cell(2, 2))
	g.Cell(2, 2).link(g.Cell(2, 3))
	//2V
	g.Cell(2, 0).link(g.Cell(3, 0))
	g.Cell(2, 2).link(g.Cell(3, 2))
	g.Cell(2, 3).link(g.Cell(3, 3))
	//3>
	g.Cell(3, 0).link(g.Cell(3, 1))
	//3V

	str := asAscii(g)

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
