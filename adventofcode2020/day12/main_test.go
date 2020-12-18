package main

import (
	"strings"
	"testing"
)

func Test_day12_task1_parse(t *testing.T) {
	is := parseInstructions(strings.NewReader("F10\nN3\nF7\nR90\nF11"))

	if is[0].action != "F" || is[0].argument != 10 {
		t.Errorf("Wrong instruction %v", is[0])
	}
	if is[3].action != "R" || is[3].argument != 90 {
		t.Errorf("Wrong instruction %v", is[3])
	}
}

func Test_day12_task1_move(t *testing.T) {
	is := parseInstructions(strings.NewReader("F10\nN3\nF7\nR90\nF11"))
	s := ship{}

	s.doAll(is)

	if s.east != 17 || s.north != -8 {
		t.Errorf("Wrong ship location %#v", s)
	}
}

func Test_day12_task2_move(t *testing.T) {
	is := parseInstructions(strings.NewReader("F10\nN3\nF7\nR90\nF11"))
	s := ship{}

	s.doAll2(is)

	if s.east != 214 || s.north != -72 {
		t.Errorf("Wrong ship location %#v", s)
	}
}
