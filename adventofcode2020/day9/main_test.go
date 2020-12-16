package main

import (
	"strings"
	"testing"
)

var testData = strings.TrimPrefix(`
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`, "\n")

func Test_day9_task1(t *testing.T) {
	n := parseNumbers(strings.NewReader(testData))

	fwn := n.firstWrongNumber(5)

	if fwn != 127 {
		t.Errorf("Wrong 'first wrong number' %d", fwn)
	}

}

func Test_day9_task2(t *testing.T) {
	n := parseNumbers(strings.NewReader(testData))

	l, h := n.findLowHigh( 127)

	if l != 15 || h != 47 {
		t.Errorf("Wrong low %d or high %d values", l, h)
	}

}
