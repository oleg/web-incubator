package mazes

import "math/rand"

func BinaryTree(grid Grid) {
	grid.EachCells(func(cell *Cell) {
		var neighbors []*Cell
		if cell.north != nil {
			neighbors = append(neighbors, cell.north)
		}
		if cell.east != nil {
			neighbors = append(neighbors, cell.east)
		}
		if len(neighbors) > 0 {
			cell.link(neighbors[rand.Intn(len(neighbors))])
		}
	})
}
