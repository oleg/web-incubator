package maze

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_new_cell(t *testing.T) {
	cell := NewCell(0, 0)

	assert.Equal(t, 0, cell.row)
	assert.Equal(t, 0, cell.column)
	assert.Empty(t, cell.links)
}

func Test_linked__two_cells_unlinked_if_they_are_not_linked(t *testing.T) {
	c1 := NewCell(0, 0)
	c2 := NewCell(0, 1)

	assert.False(t, c1.Linked(c2))
	assert.False(t, c2.Linked(c1))
}

func Test_linked__both_cells_are_linked_after_one_linked_to_another(t *testing.T) {
	c1 := NewCell(0, 0)
	c2 := NewCell(0, 1)

	c1.Link(c2)

	assert.True(t, c1.Linked(c2))
	assert.True(t, c2.Linked(c1))
}
