package main

import (
	"testing"
)

func Test_day17_task1_count(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"case 1", ".#.\n..#\n###", 5},
		{"case 2", ".#.\n..#\n#..", 3},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := newUni(test.input)

			active := s.countActive()

			if active != test.expected {
				t.Errorf("Wrong number of active cubes %v, expected %v", active, test.expected)
			}
		})
	}
}

func Test_day17_task1_countActiveNeighbors(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		input    string
		x, y, z  int
	}{
		{"case 1", 3, ".##\n..#\n...", 0, 1, 1},
		{"case 2", 1, ".##\n..#\n...", 0, 0, 0},
		{"case 3", 0, ".##\n..#\n...", 0, 2, 0},
		{"case 4", 0, ".##\n..#\n#..", 0, 2, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := newUni(test.input)

			active := s.countActiveNeighbors(test.x, test.y, test.z)

			if active != test.expected {
				t.Errorf("Wrong number of active cubes %v, expected %v", active, test.expected)
			}
		})
	}
}

func Test_day17_task1_grow(t *testing.T) {
	s := newUni(".#.\n..#\n###")

	s.grow()

	active := s.countActive()
	if active != 11 {
		t.Errorf("Wrong number of active cubes %v, expected %v", active, 11)
	}

	s.grow()
	s.grow()
	s.grow()
	s.grow()
	s.grow()

	active = s.countActive()
	if active != 112 {
		t.Errorf("Wrong number of active cubes %v, expected %v", active, 112)
	}
}