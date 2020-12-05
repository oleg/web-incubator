package generator

import (
	"math/rand"
	"mazes/maze"
)

func BinaryTree(grid *maze.Grid) {
	grid.EachCells(func(cell *maze.Cell) {
		if neighbor := chooseNeighbor(grid, cell); neighbor != nil {
			cell.Link(neighbor)
		}
	})
}

func chooseNeighbor(grid *maze.Grid, cell *maze.Cell) *maze.Cell {
	north := grid.North(cell)
	east := grid.East(cell)

	if north != nil && east != nil {
		if rand.Intn(2) == 0 {
			return north
		}
		return east
	}
	if north != nil {
		return north
	}
	if east != nil {
		return east
	}
	return nil
}
