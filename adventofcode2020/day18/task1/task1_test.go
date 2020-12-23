package task1

import (
	"testing"
)

func Test_day18_task1(t *testing.T) {
	tests := []struct {
		expr     string
		expected int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}
	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {

			res := Evaluate(test.expr)

			if res != test.expected {
				t.Errorf("Unexpected evaluation result %v expected %v", res, test.expected)
			}
		})
	}
}
