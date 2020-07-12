package mazes_in_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_new_cell(t *testing.T) {
	cell := NewCell(0, 0)

	assert.Equal(t, 0, cell.row)
	assert.Equal(t, 0, cell.column)
	assert.Empty(t, cell.links)
	assert.Nil(t, cell.east)
	assert.Nil(t, cell.north)
	assert.Nil(t, cell.west)
	assert.Nil(t, cell.south)
}

func Test_is_linked_two_not_linked(t *testing.T) {
	c1 := NewCell(0, 0)
	c2 := NewCell(0, 1)

	assert.False(t, c1.linked(c2))
	assert.False(t, c2.linked(c1))
}

func Test_is_linked_two_both_linked(t *testing.T) {
	c1 := NewCell(0, 0)
	c2 := NewCell(0, 1)

	c1.link(c2)

	assert.True(t, c1.linked(c2))
	assert.True(t, c2.linked(c1))
}

func Test_is_linked_two_both_linked(t *testing.T) {
	c1 := NewCell(0, 0)
	c2 := NewCell(0, 1)

	c1.link(c2)

	assert.True(t, c1.linked(c2))
	assert.True(t, c2.linked(c1))
}
