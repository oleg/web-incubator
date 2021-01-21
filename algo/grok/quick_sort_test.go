package grok

import (
	"algo/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_QuickSortNew(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			res := QuickSortNew(test.Array)
			assert.Equal(t, test.Expected, res, test.Name)
		})
	}
}

func Test_QuickSortLomuto(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			QuickSortLomuto(test.Array)
			assert.Equal(t, test.Expected, test.Array, test.Name)
		})
	}
}
