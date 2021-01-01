package main

import (
	"testing"
)

func Test_day23_task1_destination(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		cups     []int
	}{
		{"simple case", 4, []int{5, 1, 2, 3, 4}},
		{"picked up", 5, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{"picked up 2", 5, []int{7, 1, 6, 2, 3, 4, 5}},
		{"wraps around", 7, []int{4, 1, 2, 3, 5, 6, 7}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cs := newCycleBuffer(test.cups)
			destination := cs.dest(cs.cut3())

			if destination != test.expected {
				t.Errorf("Wrong destination %v, expected %v", destination, test.expected)
			}
		})
	}
}

func Test_day23_task1_move(t *testing.T) {
	cb := newCycleBuffer([]int{3, 8, 9, 1, 2, 5, 4, 6, 7})
	if cb.String() != "389125467" {
		t.Errorf("Wrong move: %v", cb.String())
	}
	move(cb)
	if cb.String() != "289154673" {
		t.Errorf("Wrong move: %v", cb.String())
	}
	move(cb)
	if cb.String() != "546789132" {
		t.Errorf("Wrong move: %v", cb.String())
	}
	move(cb)
	if cb.String() != "891346725" {
		t.Errorf("Wrong move: %v", cb.String())
	}
	move(cb)
	if cb.String() != "467913258" {
		t.Errorf("Wrong move: %v", cb.String())
	}
	move(cb)
	if cb.String() != "136792584" {
		t.Errorf("Wrong move: %v", cb.String())
	}
	move(cb)
	if cb.String() != "936725841" {
		t.Errorf("Wrong move: %v", cb.String())
	}
	move(cb)
	if cb.String() != "258367419" {
		t.Errorf("Wrong move: %v", cb.String())
	}
	move(cb)
	if cb.String() != "674158392" {
		t.Errorf("Wrong move: %v", cb.String())
	}
	move(cb)
	if cb.String() != "574183926" {
		t.Errorf("Wrong move: %v", cb.String())
	}
	move(cb)
	if cb.String() != "837419265" {
		t.Errorf("Wrong move: %v", cb.String())
	}
}

func Test_day23_task1_move10(t *testing.T) {
	cb := newCycleBuffer([]int{3, 8, 9, 1, 2, 5, 4, 6, 7})

	moveN(10, cb)

	if cb.String() != "837419265" {
		t.Errorf("Wrong move: %v", cb.String())
	}
}

func Test_day23_task1_move100(t *testing.T) {
	cb := newCycleBuffer([]int{3, 8, 9, 1, 2, 5, 4, 6, 7})

	moveN(100, cb)

	if cb.String() != "167384529" {
		t.Errorf("Wrong move: %v", cb.String())
	}
}
