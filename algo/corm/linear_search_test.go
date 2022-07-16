package corm

import (
	"algo/assert"
	"testing"
)

func Test_LinearSearch(t *testing.T) {
	type expected struct {
		found bool
		index int
	}
	tests := []struct {
		name     string
		arr      []int
		val      int
		expected expected
	}{
		{"empty - not found", []int{}, 10, expected{false, 0}},
		{"1 elem - found", []int{10}, 10, expected{true, 0}},
		{"5 elem - found", []int{10, 3, 5, -9, 100}, 5, expected{true, 2}},
		{"5 elem - not found", []int{10, 3, 5, -9, 100}, 200, expected{false, 0}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			found, index := LinearSearch(test.arr, test.val)

			assert.Equal(t, found, test.expected.found)
			assert.Equal(t, index, test.expected.index)
		})
	}
}
