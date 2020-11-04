package mazes

import "math/rand"

func Sidewinder(grid Grid) {
	for _, row := range grid.cells {
		var run []*Cell
		for _, cell := range row {
			run = append(run, cell)

			shouldCloseOut := cell.east == nil ||
				(cell.north != nil && rand.Intn(2) == 0)

			if shouldCloseOut {
				member := run[rand.Intn(len(run))]
				if member.north != nil {
					member.link(member.north)
				}
				run = []*Cell{}
			} else {
				cell.link(cell.east)
			}
		}

	}
}
