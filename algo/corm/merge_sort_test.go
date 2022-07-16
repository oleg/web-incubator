package corm

import (
	"algo/assert"
	"algo/data"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			res := MergeSort(test.Array)
			assert.EqualSlice(t, res, test.Expected)
		})
	}
}
