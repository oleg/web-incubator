package main

import (
	"github.com/oleg/incubator/adventofcode2020/misc"
	"strings"
)

func main() {
	text := misc.MustReadFileToString("day3/input.txt")
	forest := NewForest(text)
	println(forest.CountTrees(1, 3))
	println(forest.CountProduct())
}

type Forest struct {
	width int
	lines []string
}

func NewForest(text string) Forest {
	lines := strings.Split(text, "\n")
	return Forest{width: len(lines[0]), lines: lines}
}

func (f *Forest) Get(row, column int) string {
	return string(f.lines[row][column%f.width])
}
func (f *Forest) CountTrees(rowInc, columnInc int) int {
	count := 0
	row := 0
	column := 0
	for row < len(f.lines)-1 {
		row += rowInc
		column += columnInc
		if f.Get(row, column) == "#" {
			count++
		}
	}
	return count
}

func (f *Forest) CountProduct() int {
	return f.CountTrees(1, 1) *
		f.CountTrees(1, 3) *
		f.CountTrees(1, 5) *
		f.CountTrees(1, 7) *
		f.CountTrees(2, 1)
}
