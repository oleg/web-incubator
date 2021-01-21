package grok

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BinarySearch(t *testing.T) {
	type expected struct {
		found bool
		index int
	}
	tests := []struct {
		name     string
		expected expected
		val      int
		arr      []int
	}{
		{"empty - not found", expected{false, 0}, 10, []int{}},
		{"1 elem - found", expected{true, 0}, 10, []int{10}},
		{"5 elem - not found", expected{false, 0}, 200, []int{-9, 3, 5, 10, 100}},
		{"5 elem - found", expected{true, 2}, 5, []int{-9, 3, 5, 10, 100}},
		{"5 elem(100) - found", expected{true, 4}, 100, []int{-9, 3, 5, 10, 100}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			index, found := Search(test.val, test.arr)

			assert.Equal(t, test.expected.index, index)
			assert.Equal(t, test.expected.found, found)
		})
	}
}
