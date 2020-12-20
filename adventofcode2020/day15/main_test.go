package main

import "testing"

func Test_day15_task1(t *testing.T) {
	//expected := []int{0, 3, 6, 0, 3, 3, 1, 0, 4, 0}
	g := newGenerator(0, 3, 6)

	val := g.next()
	if val != 0 || g.step != 1 || g.last != 0 || g.stats[0].beforeLast != 1 {
		t.Errorf("Wrong generator value %v or state %#v", val, g)
	}
	val = g.next()
	if val != 3 || g.step != 2 || g.last != 3 || g.stats[3].beforeLast != 2 {
		t.Errorf("Wrong generator value %v or state %#v", val, g)
	}
	val = g.next()
	if val != 6 || g.step != 3 || g.last != 6 || g.stats[6].beforeLast != 3 {
		t.Errorf("Wrong generator value %v or state %#v", val, g)
	}
	val = g.next()
	if val != 0 || g.step != 4 || g.last != 0 || g.stats[0].beforeLast != 1 || g.stats[0].last != 4 {
		t.Errorf("Wrong generator value %v or state %#v", val, g)
	}
	val = g.next()
	if val != 3 || g.step != 5 || g.last != 3 || g.stats[3].beforeLast != 2 || g.stats[3].last != 5 {
		t.Errorf("Wrong generator value %v or state %#v", val, g)
	}
	val = g.next()
	if val != 3 || g.step != 6 || g.last != 3 || g.stats[3].beforeLast != 5 || g.stats[3].last != 6 {
		t.Errorf("Wrong generator value %v or state %#v", val, g)
	}
	val = g.next()
	if val != 1 || g.step != 7 || g.last != 1 || g.stats[1].beforeLast != 7 || g.stats[1].last != 7 {
		t.Errorf("Wrong generator value %v or state %#v", val, g)
	}
}

func Test_day15_task1_all(t *testing.T) {
	g := newGenerator(0, 3, 6)
	for i, exp := range []int{0, 3, 6, 0, 3, 3, 1, 0, 4, 0} {
		val := g.next()
		if val != exp {
			t.Errorf("Wrong %vth value %v expected %v", i, val, exp)
		}
	}
}

func Test_day15_task1_count(t *testing.T) {
	g := newGenerator(0, 3, 6)

	g.skip(2020)

	if g.last != 436 {
		t.Errorf("Wrong last value %v", g.last)
	}
}