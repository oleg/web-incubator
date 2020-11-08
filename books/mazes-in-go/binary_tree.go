package mazes

import "math/rand"

func BinaryTree(grid *Grid) {
	grid.EachCells(func(cell *Cell) {
		if neighbor := chooseNeighbor(grid, cell); neighbor != nil {
			cell.link(neighbor)
		}
	})
}

func chooseNeighbor(grid *Grid, cell *Cell) *Cell {
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
