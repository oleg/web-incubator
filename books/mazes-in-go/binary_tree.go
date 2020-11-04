package mazes

import "math/rand"

func BinaryTree(grid Grid) {
	grid.EachCells(func(cell *Cell) {
		if neighbor := chooseNeighbor(cell); neighbor != nil {
			cell.link(neighbor)
		}
	})
}

func chooseNeighbor(cell *Cell) *Cell {
	if cell.north != nil && cell.east != nil {
		if rand.Intn(2) == 0 {
			return cell.north
		}
		return cell.east
	}
	if cell.north != nil {
		return cell.north
	}
	if cell.east != nil {
		return cell.east
	}
	return nil
}
