package main

import "testing"

func Test(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}

	a, b := find1(input)

	if a != 1721 {
		t.Errorf("Wrong fist number %d", a)
	}
	if b != 299 {
		t.Errorf("Wrong second number %d", b)
	}
	if (a * b) != 514579 {
		t.Errorf("Wrong solution %d", a*b)
	}
}
