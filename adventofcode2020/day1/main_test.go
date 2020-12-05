package main

import "testing"

func Test_day1_task1(t *testing.T) {
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

func Test_day1_task2(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}

	a, b, c := find2(input)

	if a != 979 {
		t.Errorf("Wrong fist number %d", a)
	}
	if b != 366 {
		t.Errorf("Wrong second number %d", b)
	}
	if c != 675 {
		t.Errorf("Wrong third number %d", c)
	}
	if (a * b * c) != 241861950 {
		t.Errorf("Wrong solution %d", a*b*c)
	}
}
