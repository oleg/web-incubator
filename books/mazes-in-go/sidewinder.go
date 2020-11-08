package mazes

import "math/rand"

func Sidewinder(grid *Grid) {
	for _, row := range grid.cells {
		var run []*Cell
		for _, cell := range row {
			run = append(run, cell)

			shouldCloseOut := grid.East(cell) == nil ||
				(grid.North(cell) != nil && rand.Intn(2) == 0)

			if shouldCloseOut {
				member := run[rand.Intn(len(run))]
				north := grid.North(member)
				if north != nil {
					member.link(north)
				}
				run = []*Cell{}
			} else {
				cell.link(grid.East(cell))
			}
		}

	}
}
