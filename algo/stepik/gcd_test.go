package stepik

import (
	"testing"
)

func Test_gcd(t *testing.T) {
	tests := []struct {
		name           string
		a, b, expected int
	}{
		{"18 35", 18, 35, 1},
		{"14159572 63967072", 14159572, 63967072, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			res := gcd(test.a, test.b)

			if res != test.expected {
				t.Errorf("wrong gcd %d", res)
			}
		})
	}
}
