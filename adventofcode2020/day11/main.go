package main

import (
	"bufio"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
)

func main() {
	a := parseArea(misc.MustOpen("day11/input.txt"))
	stabilize(a)
	println(a.countOccupied())
	b := parseArea(misc.MustOpen("day11/input.txt"))
	stabilize2(b)
	println(b.countOccupied())
}

type area [][]string

func stabilize(a *area) {
	b := newArea(a)
	for !a.eq(b) {
		a.move(b)
		b, a = a, b
	}
}
func stabilize2(a *area) {
	b := newArea(a)
	for !a.eq(b) {
		a.move2(b)
		b, a = a, b
	}
}

func newArea(a *area) *area {
	na := make([][]string, len(*a))
	for i, r := range *a {
		na[i] = make([]string, len(r))
	}
	b := area(na)
	return &b
}

func parseArea(reader io.Reader) *area {
	scanner := bufio.NewScanner(reader)
	seats := area(make([][]string, 0))
	for scanner.Scan() {
		text := scanner.Text()
		line := make([]string, 0, len(text))
		for _, v := range text {
			line = append(line, string(v))
		}
		seats = append(seats, line)
	}
	return &seats
}

func (a *area) move(b *area) {
	bb := *b
	for i, r := range *a {
		for j, v := range r {
			bb[i][j] = transform(v, 4, a.countOccupiedAt(i, j))
		}
	}
}
func (a *area) move2(b *area) {
	bb := *b
	for i, r := range *a {
		for j, v := range r {
			bb[i][j] = transform(v, 5, a.countOccupiedAt2(i, j))
		}
	}
}

func transform(v string, max, occupiedCount int) string {
	if v == "L" && occupiedCount == 0 {
		return "#"
	}
	if v == "#" && occupiedCount >= max {
		return "L"
	}
	return v
}

func (a *area) countOccupiedAt(row, column int) int {
	count := 0
	for i := row - 1; i <= row+1; i++ {
		for j := column - 1; j <= column+1; j++ {
			if (i != row || j != column) && a.itemAt(i, j) == "#" {
				count++
			}
		}
	}
	return count
}
func (a *area) countOccupiedAt2(row, column int) int {
	count := 0
	count += funcName1(row, column, a)
	count += funcName2(row, column, a)
	count += funcName3(row, column, a)
	count += funcName4(row, column, a)
	count += funcName5(row, column, a)
	count += funcName6(row, column, a)
	count += funcName7(row, column, a)
	count += funcName8(row, column, a)
	return count
}

func funcName8(row int, column int, a *area) int {
	aa := *a
	i := row + 1
	j := column - 1
	for ; i < len(aa) && j >= 0; {
		if aa[i][j] == "#" {
			return 1
		}
		if aa[i][j] == "L" {
			return 0
		}
		i++
		j--
	}
	return 0
}

func funcName7(row int, column int, a *area) int {
	aa := *a
	i := row - 1
	j := column + 1
	for ; i >= 0 && j < len(aa[0]); {
		if aa[i][j] == "#" {
			return 1
		}
		if aa[i][j] == "L" {
			return 0
		}
		i--
		j++
	}
	return 0
}

func funcName6(row int, column int, a *area) int {
	aa := *a
	i := row - 1
	j := column - 1
	for ; i >= 0 && j >= 0; {
		if aa[i][j] == "#" {
			return 1
		}
		if aa[i][j] == "L" {
			return 0
		}
		i--
		j--
	}
	return 0
}

func funcName5(row int, column int, a *area) int {
	aa := *a
	i := row + 1
	j := column + 1
	for ; i < len(aa) && j < len(aa[0]); {
		if aa[i][j] == "#" {
			return 1
		}
		if aa[i][j] == "L" {
			return 0
		}
		i++
		j++
	}
	return 0
}

func funcName4(row int, column int, a *area) int {
	aa := *a
	for i := column - 1; i >= 0; i-- {
		if aa[row][i] == "#" {
			return 1
		}
		if aa[row][i] == "L" {
			return 0
		}
	}
	return 0
}

func funcName3(row int, column int, a *area) int {
	aa := *a
	for i := column + 1; i < len(aa[0]); i++ {
		if aa[row][i] == "#" {
			return 1
		}
		if aa[row][i] == "L" {
			return 0
		}
	}
	return 0
}

func funcName2(row int, column int, a *area) int {
	aa := *a
	for i := row - 1; i >= 0; i-- {
		if aa[i][column] == "#" {
			return 1
		}
		if aa[i][column] == "L" {
			return 0
		}
	}
	return 0
}

func funcName1(row int, column int, a *area) int {
	aa := *a
	for i := row + 1; i < len(aa); i++ {
		if aa[i][column] == "#" {
			return 1
		}
		if aa[i][column] == "L" {
			return 0
		}
	}
	return 0
}

func (a *area) countOccupied() int {
	count := 0
	for _, r := range *a {
		for _, v := range r {
			if v == "#" {
				count++
			}
		}
	}
	return count
}

func (a *area) itemAt(row, column int) string {
	if row < 0 || row >= len(*a) {
		return ""
	}
	if column < 0 || column >= len((*a)[row]) {
		return ""
	}
	return (*a)[row][column]
}

func (a *area) eq(b *area) bool {
	aa := *a
	bb := *b
	for i := 0; i < len(aa); i++ {
		for j := 0; j < len(aa[i]); j++ {
			if aa[i][j] != bb[i][j] {
				return false
			}
		}
	}
	return true
}
