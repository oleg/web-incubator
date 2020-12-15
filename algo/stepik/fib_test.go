package stepik

import (
	"fmt"
	"testing"
)

func Test_fib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		fib  int
	}{
		{"0", 0, 0},
		{"1", 1, 1},
		{"2", 2, 1},
		{"3", 3, 2},
		{"4", 4, 3},
		{"5", 5, 5},
		{"6", 6, 8},
		{"7", 7, 13},
		{"8", 8, 21},
		{"9", 9, 34},
		{"10", 10, 55},
		{"30", 30, 832040},
		{"50", 50, 12586269025},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fib := fib(test.n)

			if fib != test.fib {
				t.Errorf("Wrong fib value %d, expected %d", fib, test.fib)
			}

		})
	}
}

func Test_fibLd(t *testing.T) {
	tests := []struct {
		name string
		n    int
		fib  int
	}{
		{"0", 0, 0},
		{"1", 1, 1},
		{"2", 2, 1},
		{"3", 3, 2},
		{"4", 4, 3},
		{"5", 5, 5},
		{"6", 6, 8},
		{"7", 7, 3},
		{"8", 8, 1},
		{"9", 9, 4},
		{"10", 10, 5},
		{"30", 30, 0},
		{"50", 50, 5},
		{"137", 137, 7},
		{"293", 293, 3},
		{"299", 299, 1},
		{"300", 300, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fib := fibLd(test.n)

			if fib != test.fib {
				t.Errorf("Wrong fib value %d, expected %d", fib, test.fib)
			}

		})
	}
}

func Test_fibBigMod(t *testing.T) {
	tests := []struct {
		name string
		n, m int
		fib  int
	}{
		{"0 4", 0, 4, 0},
		{"1 4", 1, 4, 1},
		{"9 2", 9, 2, 0},
		{"10 2", 10, 2, 1},
		{"10 4", 10, 4, 3},
		{"1025 55", 1025, 55, 5},
		{"12589 369", 12589, 369, 89},
		{"1598753 25897", 1598753, 25897, 20305},
		{"60282445765134413 2263", 60282445765134413, 2263, 974},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fib := fibBigMod(test.n, test.m)
			fmt.Println(fib)
			if fib != test.fib {
				t.Errorf("Wrong fib value %d, expected %d", fib, test.fib)
			}

		})
	}
}
