package main

func main() {
	g1 := newGenerator(20, 0, 1, 11, 6, 3)
	g1.skip(2020)
	println(g1.last)

	g2 := newGenerator(20, 0, 1, 11, 6, 3)
	g2.skip(30000000)
	println(g2.last)
}

type steps struct {
	beforeLast, last int
}
type generator struct {
	init  []int
	stats map[int]steps
	step  int
	last  int
}

func newGenerator(init ...int) *generator {
	return &generator{init: init, stats: make(map[int]steps)}
}

func (g *generator) next() int {
	var curr int
	if g.step < len(g.init) {
		curr = g.init[g.step]
	} else {
		stat := g.stats[g.last]
		curr = stat.last - stat.beforeLast
	}
	g.step++
	stat, found := g.stats[curr]
	if found {
		stat.beforeLast, stat.last = stat.last, g.step
	} else {
		stat.beforeLast, stat.last = g.step, g.step
	}
	g.stats[curr] = stat
	g.last = curr
	return curr
}

func (g *generator) skip(index int) {
	for i := 0; i < index; i++ {
		g.next()
	}
}
