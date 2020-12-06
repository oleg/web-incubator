package main

import (
	"strings"
	"testing"
)

var testForestRaw = strings.Trim(`
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`, "\n")

func Test_day3_task1_new_forest(t *testing.T) {
	forest := NewForest(testForestRaw)

	if len(forest.lines) != 11 {
		t.Errorf("Wrong height %d", len(forest.lines))
	}

	if forest.width != 11 {
		t.Errorf("Wrong height %d", forest.width)
	}
}

func Test_day3_task1_index_forest(t *testing.T) {
	tests := []struct {
		name         string
		row, columnt int
		expected     string
	}{
		{"case 1", 0, 0, "."},
		{"case 2", 0, 3, "#"},
		{"case 3", 1, 3, "."},
		{"case 4", 2, 6, "#"},
		{"case 5", 0, 11, "."},
		{"case 6", 10, 15, "#"},
	}
	forest := NewForest(testForestRaw)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			e := forest.Get(test.row, test.columnt)

			if e != test.expected {
				t.Errorf("Wrong element at (%d,%d): %s", test.row, test.columnt, e)
			}
		})
	}
}

func Test_day3_task1_count_trees(t *testing.T) {
	forest := NewForest(testForestRaw)

	count := forest.CountTrees(1, 3)

	if count != 7 {
		t.Errorf("Wrong number of trees %d", count)
	}
}


