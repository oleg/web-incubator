package main

import (
	"github.com/oleg/incubator/adventofcode2020/misc"
	"strings"
	"testing"
)

func Test_day20_parse(t *testing.T) {
	tiles := parseTiles(misc.MustOpen("testinput.txt"))

	if len(tiles) != 9 {
		t.Errorf("Wrong number of tiles %v", len(tiles))
	}
}

var tile2311 = misc.TrimNewLine(`
Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

`)

func Test_day20_print(t *testing.T) {
	tl := parseTiles(strings.NewReader(tile2311))[0]

	s := tl.String()
	if s != misc.TrimNewLine(`
n:0011010010
e:0001011001
s:1110011100
w:0100111110`) {
		t.Errorf("Wrong tile \n%v\n", s)
	}
}

func Test_day20_rotate4(t *testing.T) {
	tl := parseTiles(strings.NewReader(tile2311))[0]

	if tl.state != 0 || tl.String() != misc.TrimNewLine(`
n:0011010010
e:0001011001
s:1110011100
w:0100111110`) {
		t.Errorf("Wrong state %v or tile \n%v\n", tl.state, tl.String())
	}

	tl.change()
	if tl.state != 1 || tl.String() != misc.TrimNewLine(`
n:0100111110
e:0011010010
s:0001011001
w:1110011100`) {
		t.Errorf("Wrong state %v or tile \n%v\n", tl.state, tl.String())
	}

	tl.change()
	if tl.state != 2 || tl.String() != misc.TrimNewLine(`
n:1110011100
e:0100111110
s:0011010010
w:0001011001`) {
		t.Errorf("Wrong state %v or tile \n%v\n", tl.state, tl.String())
	}

	tl.change()
	if tl.state != 3 || tl.String() != misc.TrimNewLine(`
n:0001011001
e:1110011100
s:0100111110
w:0011010010`) {
		t.Errorf("Wrong state %v or tile \n%v\n", tl.state, tl.String())
	}
}

func Test_day20_rotate8(t *testing.T) {
	tl := parseTiles(strings.NewReader(tile2311))[0]
	tl.change()
	tl.change()
	tl.change()

	tl.change()
	if tl.state != 4 || tl.String() != misc.TrimNewLine(`
n:0100101100
e:0111110010
s:0011100111
w:1001101000`) {
		t.Errorf("Wrong state %v or tile \n%v\n", tl.state, tl.String())
	}

	tl.change()
	if tl.state != 5 || tl.String() != misc.TrimNewLine(`
n:1001101000
e:0100101100
s:0111110010
w:0011100111`) {
		t.Errorf("Wrong state %v or tile \n%v\n", tl.state, tl.String())
	}

	tl.change()
	if tl.state != 6 || tl.String() != misc.TrimNewLine(`
n:0011100111
e:1001101000
s:0100101100
w:0111110010`) {
		t.Errorf("Wrong state %v or tile \n%v\n", tl.state, tl.String())
	}

	tl.change()
	if tl.state != 7 || tl.String() != misc.TrimNewLine(`
n:0111110010
e:0011100111
s:1001101000
w:0100101100`) {
		t.Errorf("Wrong state %v or tile \n%v\n", tl.state, tl.String())
	}
}

func Test_day20_rotate9(t *testing.T) {
	original := parseTiles(strings.NewReader(tile2311))[0]
	tl := parseTiles(strings.NewReader(tile2311))[0]

	for i := 0; i < 8; i++ {
		tl.change()
	}

	if tl.normal != original.normal {
		t.Errorf("wrong tile \n%v\n expected \n%v\n", tl.normal, original.normal)
	}
	if tl.flipped != original.flipped {
		t.Errorf("wrong tile \n%v\n expected \n%v\n", tl.flipped, original.flipped)
	}
}

func Test_day20_match_north_and_south(t *testing.T) {
	tiles := parseTiles(strings.NewReader(misc.TrimNewLine(`
Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

`)))
	tUp := &tiles[0]
	tDown := &tiles[1]

	if !tDown.matchNorth(tUp) {
		t.Errorf("expected to be matched on north \n%v\n expected \n%v\n", tUp, tDown)
	}
	if !tUp.matchSouth(tDown) {
		t.Errorf("expected to be matched on south \n%v\n expected \n%v\n", tUp, tDown)
	}
}
func Test_day20_match_east_and_west(t *testing.T) {
	tiles := parseTiles(strings.NewReader(misc.TrimNewLine(`
Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

`)))
	tLeft := &tiles[0]
	tRight := &tiles[1]

	if !tLeft.matchEast(tRight) {
		t.Errorf("expected to be matched on east \n%v\n expected \n%v\n", tLeft, tRight)
	}
	if !tRight.matchWest(tLeft) {
		t.Errorf("expected to be matched on west \n%v\n expected \n%v\n", tLeft, tRight)
	}
}

func Test_day20_find_corners(t *testing.T) {
	tiles := parseTiles(misc.MustOpen("testinput.txt"))

	ids := findCorners(tiles)

	if _, ok := ids[1951]; !ok {
		t.Errorf("Wrong tiles %v", ids)
	}
	if _, ok := ids[3079]; !ok {
		t.Errorf("Wrong tiles %v", ids)
	}
	if _, ok := ids[1171]; !ok {
		t.Errorf("Wrong tiles %v", ids)
	}
	if _, ok := ids[2971]; !ok {
		t.Errorf("Wrong tiles %v", ids)
	}
}

func Test_day20_task1(t *testing.T) {
	tiles := parseTiles(misc.MustOpen("testinput.txt"))

	corners := findCorners(tiles)

	if multiply(corners) != 20899048083289 {
		t.Errorf("Wrong multiplication %v", multiply(corners))
	}
}
