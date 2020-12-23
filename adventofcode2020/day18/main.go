package main

import (
	"bufio"
	"github.com/oleg/incubator/adventofcode2020/day18/task1"
	"github.com/oleg/incubator/adventofcode2020/day18/task2"
	"github.com/oleg/incubator/adventofcode2020/misc"
)

func main() {
	evaluateInput := func(f func(string) int) int {
		s := bufio.NewScanner(misc.MustOpen("day18/input.txt"))
		sum := 0
		for s.Scan() {
			sum += f(s.Text())
		}
		return sum
	}

	println(evaluateInput(task1.Evaluate))
	println(evaluateInput(task2.Evaluate))
}
