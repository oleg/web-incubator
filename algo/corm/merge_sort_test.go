package corm

import (
	"algo/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			res := MergeSort(test.Array)
			assert.Equal(t, test.Expected, res, test.Name)
		})
	}
}
