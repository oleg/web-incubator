package main

import (
	"bufio"
	"fmt"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
	"strconv"
	"strings"
)

func main() {
	println(multiply(findCorners(parseTiles(misc.MustOpen("day20/input.txt")))))
}

type sides struct {
	n, e, s, w int
}

func (s *sides) rotateLeft() {
	s.n, s.e, s.s, s.w = s.w, s.n, s.e, s.s
}

type tile struct {
	id              int
	state           byte
	normal, flipped sides
}

func (t *tile) String() string {
	return fmt.Sprintf("n:%010b\ne:%010b\ns:%010b\nw:%010b",
		t.normal.n, t.normal.e, t.normal.s, t.normal.w)
}

func (t *tile) change() {
	t.normal.rotateLeft()
	t.flipped.rotateLeft()
	if t.state == 3 || t.state == 7 {
		t.normal, t.flipped = t.flipped, t.normal
	}
	t.state++
}

func (t *tile) matchNorth(o *tile) bool {
	return t.normal.n == o.flipped.s
}
func (t *tile) matchSouth(o *tile) bool {
	return t.normal.s == o.flipped.n
}

func (t *tile) matchEast(o *tile) bool {
	return t.normal.e == o.flipped.e
}
func (t *tile) matchWest(o *tile) bool {
	return t.normal.w == o.flipped.w
}

func newTile(id int, data []string) tile {
	n := data[0]
	e := ""
	for _, v := range data {
		e += string(v[9])
	}
	s := data[9]
	w := ""
	for _, v := range data {
		w += string(v[0])
	}
	return tile{
		id: id,
		normal: sides{
			n: toNum(n, false),
			e: toNum(e, false),
			s: toNum(s, true),
			w: toNum(w, true),
		},
		flipped: sides{
			n: toNum(n, true),
			w: toNum(e, true),
			s: toNum(s, false),
			e: toNum(w, false),
		},
	}
}

func toNum(str string, reverse bool) int {
	var b strings.Builder
	for i := range str {
		if reverse {
			i = 9 - i
		}
		d := '1'
		if str[i] == '.' {
			d = '0'
		}
		b.WriteRune(d)
	}

	num, err := strconv.ParseInt(b.String(), 2, 32)
	if err != nil {
		panic(err)
	}
	return int(num)
}

func parseTiles(input io.Reader) []tile {
	scanner := bufio.NewScanner(input)
	data := make([]string, 10)
	id := 0
	line := 0
	tiles := make([]tile, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			tiles = append(tiles, newTile(id, data))
			line = 0
			continue
		}
		if strings.HasPrefix(text, "Tile") {
			id = misc.MustAtoi(text[5:9])
			continue
		}
		data[line] = text
		line++
	}
	return tiles
}

func findCorners(tiles []tile) map[int]struct{} {
	ids := make(map[int]struct{}, 0)
	seqFalse := 0
	for _, t := range tiles {
		for ti := 0; ti < 8; ti++ {
			matchedSide := false
			for _, o := range tiles {
				if t.id != o.id {
					matched := false
					for oi := 0; oi < 8; oi++ {
						matched = matched || t.matchNorth(&o)
						(&o).change()
					}
					matchedSide = matchedSide || matched
				}
			}
			if !matchedSide {
				seqFalse++
				if seqFalse == 2 {
					ids[t.id] = struct{}{}
				}
			} else {
				seqFalse = 0
			}
			t.change()
		}
	}
	return ids
}

func multiply(m map[int]struct{}) int {
	mult := 1
	for k := range m {
		mult *= k
	}
	return mult
}
