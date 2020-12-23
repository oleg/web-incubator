package task2

import (
	"testing"
)

func Test_day18_task2(t *testing.T) {
	tests := []struct {
		expr     string
		expected int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
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
