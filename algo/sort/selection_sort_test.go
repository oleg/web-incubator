package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SelectionSort(t *testing.T) {
	for _, test := range TestData {
		t.Run(test.name, func(t *testing.T) {
			SelectionSort(test.arr)
			assert.Equal(t, test.expected, test.arr, test.name)
		})
	}
}
