package corm

import (
	"algo/assert"
	"algo/data"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			InsertionSort(test.Array)
			assert.EqualSlice(t, test.Array, test.Expected)
		})
	}
}

func Test_InsertionSortReverse(t *testing.T) {
	for _, test := range data.ReverseTestData() {
		t.Run(test.Name, func(t *testing.T) {
			InsertionSortReverse(test.Array)
			assert.EqualSlice(t, test.Array, test.Expected)
		})
	}
}
