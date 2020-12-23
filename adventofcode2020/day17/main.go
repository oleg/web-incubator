package main

import (
	"github.com/oleg/incubator/adventofcode2020/misc"
	"strings"
)

func main() {
	u := newUni(misc.MustReadFileToString("day17/input.txt"))
	for i := 0; i < 6; i++ {
		u.grow()
	}
	println(u.countActive())
}

type line map[int]bool
type plane map[int]line
type space map[int]plane

type pair struct {
	from, to int
}

type uni struct {
	x, y, z pair
	space   space
}

func newUni(input string) *uni {
	lines := strings.Split(input, "\n")
	u := uni{
		x:     pair{0, 1},
		y:     pair{0, len(lines)},
		z:     pair{0, len(lines[0])},
		space: space{}}
	u.each(func(x, y, z int, v bool) {
		u.set(x, y, z, lines[y][z] == '#')
	})
	return &u
}

func (u *uni) each(do func(int, int, int, bool)) {
	for x := u.x.from; x < u.x.to; x++ {
		for y := u.y.from; y < u.y.to; y++ {
			for z := u.z.from; z < u.z.to; z++ {
				do(x, y, z, u.space[x][y][z])
			}
		}
	}

}

func (u *uni) set(x, y, z int, v bool) {
	if u.space[x] == nil {
		u.space[x] = plane{}
	}
	if u.space[x][y] == nil {
		u.space[x][y] = line{}
	}
	u.space[x][y][z] = v
}

func (u *uni) countActive() int {
	count := 0
	u.each(func(x, y, z int, v bool) {
		if v {
			count++
		}
	})
	return count
}
func (u *uni) grow() {
	u.x.from--
	u.x.to++
	u.y.from--
	u.y.to++
	u.z.from--
	u.z.to++

	o := uni{x: u.x, y: u.y, z: u.z, space: space{}}
	u.each(func(x, y, z int, v bool) {
		nbr := u.countActiveNeighbors(x, y, z)
		if v && (nbr == 2 || nbr == 3) {
			o.set(x, y, z, true)
		}
		if !v && nbr == 3 {
			o.set(x, y, z, true)
		}
	})
	u.space = o.space
}

func (u *uni) countActiveNeighbors(xx, yy, zz int) int {
	count := 0
	for x := xx - 1; x <= xx+1; x++ {
		for y := yy - 1; y <= yy+1; y++ {
			for z := zz - 1; z <= zz+1; z++ {
				if (x != xx || y != yy || z != zz) && u.space[x][y][z] {
					count++
				}
			}
		}
	}
	return count
}
