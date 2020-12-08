package main

import (
	"bufio"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"math"
)

func main() {
	reader := misc.MustOpen("day5/input.txt")
	scanner := bufio.NewScanner(reader)
	maxId := 0
	for scanner.Scan() {
		seatId := seatId(parseRowAndColumn(scanner.Text()))
		if seatId > maxId {
			maxId = seatId
		}
	}
	println(maxId)

}

func seatId(row, column int) int {
	return row*8 + column
}

func parseRowAndColumn(code string) (int, int) {
	row := search(code[:7], 'F', 'B')
	column := search(code[7:], 'L', 'R')
	return row, column
}

func search(str string, firstHalf, secondHalf rune) int {
	lower := 0
	upper := int(math.Pow(2, float64(len(str))))
	for _, v := range str {
		delta := (upper - lower) / 2
		if v == firstHalf {
			upper -= delta
		}
		if v == secondHalf {
			lower += delta
		}
	}
	return lower
}
