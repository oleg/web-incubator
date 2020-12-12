package main

import (
	"bufio"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
)

func main() {
	reader := misc.MustOpen("day6/input.txt")

	count := 0
	for _, g := range parseGroups(reader) {
		count += countUniqueAnswers(g)
	}
	println(count)
}

type group struct {
	lines []string
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
			g.lines = append(g.lines, text)
		}
	}
	groups = append(groups, g)
	return groups
}

func countUniqueAnswers(g group) int {
	answers := make(map[rune]struct{}, 0)
	for _, line := range g.lines {
		for _, r := range line {
			answers[r] = struct{}{}
		}
	}
	return len(answers)
}
func countRepeatedAnswers(g group) int {
	answers := make(map[rune]struct{}, 0)
	for _, line := range g.lines {
		for _, r := range line {
			answers[r] = struct{}{}
		}
	}
	return len(answers)
}
