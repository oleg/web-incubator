package corm

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

