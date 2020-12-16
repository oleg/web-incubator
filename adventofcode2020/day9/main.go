package main

import (
	"bufio"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
	"sort"
)

func main() {
	n := parseNumbers(misc.MustOpen("day9/input.txt"))
	println(n.firstWrongNumber(25))
	l, h := n.findLowHigh(1124361034)
	println(l + h)
}

type numbers []int

func parseNumbers(reader io.Reader) numbers {
	scanner := bufio.NewScanner(reader)
	n := numbers{}
	for scanner.Scan() {
		n = append(n, misc.MustAtoi(scanner.Text()))
	}
	return n
}

func (n numbers) firstWrongNumber(preambleLen int) int {
	for i := preambleLen; i < len(n); i++ {
		if !findSum(n[i], n[i-preambleLen:i]) {
			return n[i]
		}
	}
	return 0
}

func findSum(s int, n numbers) bool {
	for _, x := range n {
		for _, y := range n {
			if x != y && x+y == s {
				return true
			}
		}
	}
	return false
}

func (n numbers) findLowHigh(sum int) (int, int) {
	for i := 0; i < len(n); i++ {
		start := n[i]
		if toL, found := sumF(start, n[i+1:], sum); found {
			to := (i + 1) + toL
			sort.Ints(n[i : to+1])
			return n[i], n[to]
		}
	}
	return 0, 0
}

func sumF(start int, rest numbers, expected int) (int, bool) {
	s := start
	for i, v := range rest {
		s += v
		if s == expected {
			return i, true
		}
	}
	return -1, false
}
