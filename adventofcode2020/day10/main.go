package main

import (
	"fmt"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"sort"
)

func main() {
	adapters := misc.MustReadInts(misc.MustOpen("day10/input.txt"))

	d1, d3 := countDiff(adapters)
	println(d1 * d3)
	count := countArrangements(adapters)
	println(count)
}

func countDiff(adapters []int) (int, int) {
	sort.Ints(adapters)
	d1 := 0
	d3 := 0
	prev := 0
	for _, curr := range adapters {
		diff := curr - prev
		switch diff {
		case 1:
			d1++
		case 3:
			d3++
		default:
			panic(fmt.Errorf("unexpected difference %d - %d = %d", curr, prev, diff))
		}
		prev = curr
	}
	return d1, d3 + 1
}

func countArrangements(adapters []int) int {
	sort.Ints(adapters)
	return countArr(0, adapters)
}

var cache = make(map[string]int)

//todo rewrite with dynamic programming
func countArr(head int, rest []int) int {
	key := fmt.Sprintf("%d-%d", head, len(rest))
	if answer, found := cache[key]; found {
		return answer
	}
	size := 0
	if len(rest) == 0 {
		return 1
	}
	for i, next := range rest {
		if next-head > 3 {
			break
		}
		size += countArr(rest[i], rest[i+1:])
	}
	cache[key] = size
	return size
}
