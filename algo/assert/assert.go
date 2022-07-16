package assert

import (
	"testing"
)

func EqualSlice[T comparable](t *testing.T, actual, expected []T) {
	if len(actual) != len(expected) {
		t.Errorf("got %v, expected %v", actual, expected)
		return
	}

	for i, a := range actual {
		if a != expected[i] {
			t.Errorf("got %v, expected %v", actual, expected)
		}
	}
}

func Equal[T comparable](t *testing.T, actual, expected T) {
	if actual != expected {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}
