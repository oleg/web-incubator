package main

import (
	"bufio"
	"fmt"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
)

func main() {
	count := flipAndCountBlack(misc.MustOpen("day24/input.txt"))
	println(count)
}

func flipAndCountBlack(reader io.Reader) int {
	r := newRepo()
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		r.walk(scanner.Text()).flip()
	}
	return r.countBlack()
}

func splitDirs(str string) []string {
	dirs := make([]string, 0)
	pos := 0
	for pos < len(str) {
		v := str[pos : pos+1]
		if v == "s" || v == "n" {
			dirs = append(dirs, str[pos:pos+2])
			pos += 2
		} else {
			dirs = append(dirs, v)
			pos += 1
		}
	}
	return dirs
}

type tile struct {
	x, y                 int
	black                bool
	e, se, sw, w, nw, ne *tile
}

func (t *tile) flip() {
	t.black = !t.black
}

func (t *tile) set(dir string, o *tile) {
	switch dir {
	case "e":
		t.e = o
	case "se":
		t.se = o
	case "sw":
		t.sw = o
	case "w":
		t.w = o
	case "nw":
		t.nw = o
	case "ne":
		t.ne = o
	default:
		panic("Unexpected direction " + dir)
	}
}

func (t *tile) key() string {
	return key(t.x, t.y)
}

func (t *tile) xyAt(dir string) (int, int) {
	x, y := inc(dir)
	return t.x + x, t.y + y
}

type repo struct {
	id    int
	tiles map[string]*tile
}

func newRepo() repo {
	r := repo{tiles: make(map[string]*tile)}
	r.createAt(0, 0)
	return r
}

func (r *repo) createAt(x, y int) *tile {
	t := &tile{x: x, y: y}
	r.tiles[t.key()] = t
	return t
}
func (r *repo) center() *tile {
	return r.tiles[key(0, 0)]
}

var allDirs = []string{"e", "se", "sw", "w", "nw", "ne"}

func (r *repo) extend(t *tile, dirStr string) *tile {
	newTile := r.createAt(t.xyAt(dirStr))
	for _, dir := range allDirs {
		linked := r.tiles[key(newTile.xyAt(dir))]
		if linked != nil {
			newTile.set(dir, linked)
			linked.set(opposite(dir), newTile)
		}
	}
	return newTile
}

func (r *repo) walk(dirs string) *tile {
	t := r.center()
	for _, dir := range splitDirs(dirs) {
		k := key(t.xyAt(dir))
		tk, ok := r.tiles[k]
		if !ok {
			tk = r.extend(t, dir)
		}
		t = tk
	}
	return t
}

func (r *repo) countBlack() int {
	count := 0
	for _, v := range r.tiles {
		if v.black {
			count++
		}
	}
	return count
}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func inc(dir string) (int, int) {
	switch dir {
	case "e":
		return 0, 1
	case "se":
		return -1, 1
	case "sw":
		return -1, 0
	case "w":
		return 0, -1
	case "nw":
		return 1, -1
	case "ne":
		return 1, 0
	default:
		panic("Unexpected direction " + dir)
	}
}

func opposite(dir string) string {
	switch dir {
	case "e":
		return "w"
	case "se":
		return "nw"
	case "sw":
		return "ne"
	case "w":
		return "e"
	case "nw":
		return "se"
	case "ne":
		return "sw"
	default:
		panic("Unexpected direction " + dir)
	}
}
