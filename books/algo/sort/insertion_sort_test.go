package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	for _, test := range TestData {
		t.Run(test.name, func(t *testing.T) {
			InsertionSort(test.arr)
			assert.Equal(t, test.expected, test.arr, test.name)
		})
	}
}

func Test_InsertionSortReverse(t *testing.T) {
	for _, test := range ReverseTestData {
		t.Run(test.name, func(t *testing.T) {
			InsertionSortReverse(test.arr)
			assert.Equal(t, test.expected, test.arr, test.name)
		})
	}
}
