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
	tests := []struct {
		name              string
		rowInc, columnInc int
		expected          int
	}{
		{"case 1", 1, 1, 2},
		{"case 2", 1, 3, 7},
		{"case 3", 1, 5, 3},
		{"case 4", 1, 7, 4},
		{"case 5", 2, 1, 2},
	}

	forest := NewForest(testForestRaw)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			count := forest.CountTrees(test.rowInc, test.columnInc)

			if count != test.expected {
				t.Errorf("Wrong number of trees %d", count)
			}
		})
	}
}

func Test_day3_task1_count_product(t *testing.T) {
	forest := NewForest(testForestRaw)

	count := forest.CountProduct()

	if count != 336 {
		t.Errorf("Wrong product %d", count)
	}
}
