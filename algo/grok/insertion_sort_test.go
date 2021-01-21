package grok

import (
	"algo/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			InsertionSort(test.Array)
			assert.Equal(t, test.Expected, test.Array, test.Name)
		})
	}
}

func Test_InsertionSort2(t *testing.T) {
	for _, test := range data.TestData() {
		t.Run(test.Name, func(t *testing.T) {
			InsertionSort2(test.Array)
			assert.Equal(t, test.Expected, test.Array, test.Name)
		})
	}
}
