package main

import (
	"fmt"
	"strings"
	"testing"
)

var state0 = strings.TrimPrefix(`
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`, "\n")

var state1 = strings.TrimPrefix(`
#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`, "\n")

var state2 = strings.TrimPrefix(`
#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##`, "\n")

var stateLast = strings.TrimPrefix(`
#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##`, "\n")

func Test_day11_task1_parse_area(t *testing.T) {
	a := *parseArea(strings.NewReader("L.LL\nLLL.\n##L."))

	seat00 := a[0][0]
	if seat00 != "L" {
		t.Errorf("Wrong seat state at 0,0 %s", seat00)
	}
	seat13 := a[1][3]
	if seat13 != "." {
		t.Errorf("Wrong seat state at 1,3 %s", seat13)
	}
	seat21 := a[2][1]
	if seat21 != "#" {
		t.Errorf("Wrong seat state at 2,1 %s", seat21)
	}

}

func Test_day11_task1_move(t *testing.T) {
	a := parseArea(strings.NewReader(state0))
	b := newArea(a)
	a0 := parseArea(strings.NewReader(state0))
	a1 := parseArea(strings.NewReader(state1))
	a2 := parseArea(strings.NewReader(state2))

	if !a.eq(a0) {
		t.Errorf("Wrong encoding for state0 \n%v \n\nexpected\n\n%v", a, a0)
	}
	a.move(b)
	if !b.eq(a1) {
		t.Errorf("State1 is incorrect \n%v \n\nexpected\n\n%v", a, a1)
	}
	b.move(a)
	if !a.eq(a2) {
		t.Errorf("state2 is incorrect \n%v \n\nexpected\n\n%v", a, a2)
	}
}

func Test_day11_task1_stabilize(t *testing.T) {
	a := parseArea(strings.NewReader(state0))
	aLast := parseArea(strings.NewReader(stateLast))

	stabilize(a)

	if !a.eq(aLast) {
		t.Errorf("wrong state \n%v\n\n expected \n\n%v\n", a, aLast)
	}

	occupied := a.countOccupied()
	if occupied != 37 {
		t.Errorf("wrong number of occupied %d", occupied)
	}
}

func Test_day11_task2_countOccupied(t *testing.T) {
	a := parseArea(strings.NewReader(strings.TrimPrefix(`
.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`, "\n")))

	atA := a.countOccupiedAt2(4, 3)
	if atA != 8 {
		t.Errorf("wrong number of occupied %v, expected 8", atA)
	}

	fmt.Printf("%v\n", "====")

	b := parseArea(strings.NewReader(strings.TrimPrefix(`
.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`, "\n")))
	atB := b.countOccupiedAt2(3, 3)
	if atB != 0 {
		t.Errorf("wrong number of occupied %v, expected 0", atB)
	}
}
