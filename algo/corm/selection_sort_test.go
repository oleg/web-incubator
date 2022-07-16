package corm

import (
	"algo/assert"
	"algo/data"
	"testing"
)

func Test_SelectionSort(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			SelectionSort(test.Array)
			assert.EqualSlice(t, test.Array, test.Expected)
		})
	}
}
