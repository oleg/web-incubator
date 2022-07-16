package grok

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

func Test_InsertionSort2(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			InsertionSort2(test.Array)
			assert.EqualSlice(t, test.Array, test.Expected)
		})
	}
}
