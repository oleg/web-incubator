package main

import (
	"strings"
	"testing"
)

func Test_day2_parse_to_struct(t *testing.T) {
	text := `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

	data := ParsePassRules(strings.NewReader(text))
	if len(data) != 3 {
		t.Errorf("Wrong size %d", len(data))
	}
	if data[0].n1 != 1 ||
		data[0].n2 != 3 ||
		data[0].letter != "a" ||
		data[0].password != "abcde" {
		t.Errorf("Wrong struct %v", data[0])
	}
}

func Test_day2_valid1(t *testing.T) {
	tests := []struct {
		name     string
		rule     PassRule
		expected bool
	}{
		{"case 1", PassRule{1, 3, "a", "abcde"}, true},
		{"case 2", PassRule{1, 3, "b", "cdefg"}, false},
		{"case 3", PassRule{2, 9, "c", "ccccccccc"}, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			valid1 := test.rule.valid1()

			if valid1 != test.expected {
				t.Errorf("Expected %v to be valid %v", test.rule, test.expected)
			}

		})
	}
}
