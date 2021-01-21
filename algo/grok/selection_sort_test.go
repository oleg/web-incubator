package grok

import (
	"algo/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SelectionSort(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			SelectionSort(test.Array)
			assert.Equal(t, test.Expected, test.Array, test.Name)
		})
	}
}

func Test_SelectionSortNew(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			res := SelectionSortNew(test.Array)
			assert.Equal(t, test.Expected, res, test.Name)
		})
	}
}

func Test_SelectionSortNew2(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			res := SelectionSortNew2(test.Array)
			assert.Equal(t, test.Expected, res, test.Name)
		})
	}
}
