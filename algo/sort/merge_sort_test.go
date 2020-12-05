package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	for _, test := range TestData {
		t.Run(test.name, func(t *testing.T) {
			res := MergeSort(test.arr)
			assert.Equal(t, test.expected, res, test.name)
		})
	}
}
