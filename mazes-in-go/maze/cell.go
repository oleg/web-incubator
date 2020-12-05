package maze

type Cell struct {
	row, column int
	links       map[*Cell]bool
}

func NewCell(row, column int) *Cell {
	return &Cell{
		row:    row,
		column: column,
		links:  make(map[*Cell]bool),
	}
}

func (c *Cell) Linked(o *Cell) bool {
	return c.links[o]
}

func (c *Cell) Link(o *Cell) {
	c.links[o] = true
	o.links[c] = true
}
