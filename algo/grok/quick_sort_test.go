package grok

import (
	"algo/assert"
	"algo/data"
	"testing"
)

func Test_QuickSortNew(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			res := QuickSortNew(test.Array)
			assert.EqualSlice(t, res, test.Expected)
		})
	}
}

func Test_QuickSortLomuto(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			QuickSortLomuto(test.Array)
			assert.EqualSlice(t, test.Array, test.Expected)
		})
	}
}
