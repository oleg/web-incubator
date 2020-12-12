package main

import (
	"bufio"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
	"strconv"
	"strings"
)

func main() {
	reader := misc.MustOpen("day6/input.txt")

	countUnique := 0
	countRepeated := 0
	for _, g := range parseGroups(reader) {
		countUnique += countUniqueAnswers(g)
		countRepeated += countRepeatedAnswers(g)
	}
	println(countUnique)
	println(countRepeated)
}

type group struct {
	lines []line
}

type line int64

func newLine(str string) line {
	bits := [26]bool{}
	for _, r := range str {
		bits['z'-r] = true
	}
	bitStr := ""
	for _, bit := range bits {
		bitStr += bitToStr(bit)
	}
	res, err := strconv.ParseInt(bitStr, 2, 0)
	if err != nil {
		panic(err)
	}
	return line(res)
}

func bitToStr(b bool) string {
	if b {
		return "1"
	} else {
		return "0"
	}
}

func (l line) count() int {
	binary := strconv.FormatInt(int64(l), 2)
	return strings.Count(binary, "1")
}

func parseGroups(reader io.Reader) []group {
	groups := make([]group, 0)
	scanner := bufio.NewScanner(reader)
	g := group{}
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			groups = append(groups, g)
			g = group{}
		} else {
			g.lines = append(g.lines, newLine(text))
		}
	}
	groups = append(groups, g)
	return groups
}

func countUniqueAnswers(g group) int {
	res := newLine("")
	for _, line := range g.lines {
		res |= line
	}
	return res.count()
}

func countRepeatedAnswers(g group) int {
	res := newLine("abcdefghijklmnopqrstuvwxyz")
	for _, line := range g.lines {
		res &= line
	}
	return res.count()
}
