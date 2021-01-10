package stepik

import (
	"fmt"
	"io"
	"sort"
)

func process(r io.Reader, w io.Writer) {
	writePoints(w, placeDots(readIntervals(r)))
}

func readIntervals(r io.Reader) []interval {
	var n int
	fmt.Fscan(r, &n)
	intervals := make([]interval, n)
	for i := 0; i < n; i++ {
		var from, to int
		fmt.Fscan(r, &from, &to)
		intervals[i] = interval{from: from, to: to}
	}
	return intervals
}

func writePoints(w io.Writer, dots []int) {
	fmt.Fprintln(w, len(dots))
	if len(dots) == 0 {
		return
	}
	fmt.Fprint(w, dots[0])
	for _, d := range dots[1:] {
		fmt.Fprint(w, " ", d)
	}
}

func placeDots(intervals []interval) []int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].to < intervals[j].to
	})

	if len(intervals) == 0 {
		return nil
	}

	dots := make([]int, 0)
	dot := intervals[0].to
	dots = append(dots, dot)
	for _, v := range intervals[1:] {
		if !v.contains(dot) {
			dot = v.to
			dots = append(dots, dot)
		}
	}
	return dots
}

type interval struct {
	from, to int
}

func (i *interval) contains(dot int) bool {
	return i.from <= dot && dot <= i.to
}
