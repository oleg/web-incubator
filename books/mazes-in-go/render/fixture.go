package render

import "mazes/maze"

//	+---+---+---+---+
//	|               |
//	+   +   +---+   +
//	|   |       |   |
//	+---+---+   +---+
//	|               |
//	+   +---+   +   +
//	|       |   |   |
//	+---+---+---+---+
func complex4x4Grid() *maze.Grid {
	g := maze.NewGrid(4, 4)
	//0>
	g.Cell(0, 0).Link(g.Cell(0, 1))
	g.Cell(0, 1).Link(g.Cell(0, 2))
	g.Cell(0, 2).Link(g.Cell(0, 3))
	//0V
	g.Cell(0, 0).Link(g.Cell(1, 0))
	g.Cell(0, 1).Link(g.Cell(1, 1))
	g.Cell(0, 3).Link(g.Cell(1, 3))
	//1>
	g.Cell(1, 1).Link(g.Cell(1, 2))
	//1V
	g.Cell(1, 2).Link(g.Cell(2, 2))
	//2>
	g.Cell(2, 0).Link(g.Cell(2, 1))
	g.Cell(2, 1).Link(g.Cell(2, 2))
	g.Cell(2, 2).Link(g.Cell(2, 3))
	//2V
	g.Cell(2, 0).Link(g.Cell(3, 0))
	g.Cell(2, 2).Link(g.Cell(3, 2))
	g.Cell(2, 3).Link(g.Cell(3, 3))
	//3>
	g.Cell(3, 0).Link(g.Cell(3, 1))
	//3V
	return g
}
