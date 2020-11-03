package mazes

type Cell struct {
	row, column              int
	north, east, south, west *Cell //todo make an array?
	links                    map[*Cell]bool//todo rename merged, connected
}

func NewCell(row, column int) *Cell {
	return &Cell{
		row:    row,
		column: column,
		links:  make(map[*Cell]bool),
	}
}

func (c *Cell) linked(o *Cell) bool {
	return c.links[o]
}

func (c *Cell) link(o *Cell) {
	c.links[o] = true
	o.links[c] = true
}
