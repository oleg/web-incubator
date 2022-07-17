package grok

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

func Test_SelectionSortNew(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			res := SelectionSortNew(test.Array)
			assert.EqualSlice(t, res, test.Expected)
		})
	}
}

func Test_SelectionSortNew2(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			res := SelectionSortNew2(test.Array)
			assert.EqualSlice(t, res, test.Expected)
		})
	}
}
