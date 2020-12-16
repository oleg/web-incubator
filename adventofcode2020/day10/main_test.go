package main

import (
	"github.com/oleg/incubator/adventofcode2020/misc"
	"strings"
	"testing"
)

func Test_day10_task1(t *testing.T) {
	tests := []struct {
		name                   string
		expectedD1, expectedD3 int
		data                   string
	}{
		{"small", 7, 5, "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"},
		{"big", 22, 10, "28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			adapters := misc.MustReadInts(strings.NewReader(test.data))

			d1, d3 := countDiff(adapters)

			if d1 != test.expectedD1 || d3 != test.expectedD3 {
				t.Errorf("Unexpected jolt diff %d, %d ", d1, d3)
			}
		})
	}
}
func Test_day10_task2(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		data     string
	}{
		{"one item", 1, "3"},
		{"two items", 2, "1\n2"},
		{"three items", 4, "1\n2\n3"},
		{"four items", 7, "1\n2\n3\n4"},
		{"small", 8, "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"},
		{"big", 19208, "28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			adapters := misc.MustReadInts(strings.NewReader(test.data))

			count := countArrangements(adapters)

			if count != test.expected {
				t.Errorf("Unexpected arrangements count %d, expected %d", count, test.expected)
			}
		})
	}
}


