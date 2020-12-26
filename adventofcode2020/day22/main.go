package main

import (
	"bufio"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
	"strings"
)

func main() {
	g := parseGame(misc.MustOpen("day22/input.txt"))
	g.moveToEnd()
	println(g.count())
}

type player []int

func (p *player) take() int {
	var card int
	card, *p = (*p)[0], (*p)[1:]
	return card
}

func (p *player) put(a, b int) {
	*p = append(*p, a, b)
}

func (p *player) count() int {
	count := 0
	l := len(*p)
	for i, v := range *p {
		count += (l - i) * v
	}
	return count
}

type game struct {
	player1, player2 player
}

func (g *game) move() {
	c1 := g.player1.take()
	c2 := g.player2.take()
	if c1 == c2 {
		panic("unexpected state")
	}
	if c1 > c2 {
		g.player1.put(c1, c2)
	}
	if c1 < c2 {
		g.player2.put(c2, c1)
	}
}

func (g *game) moveToEnd() {
	for len(g.player1) > 0 && len(g.player2) > 0 {
		g.move()
	}
}

func (g *game) count() int {
	if len(g.player1) == 0 {
		return g.player2.count()
	}
	if len(g.player2) == 0 {
		return g.player1.count()
	}
	panic("unexpected state")
}
func parseGame(reader io.Reader) game {
	scanner := bufio.NewScanner(reader)
	p1 := scanPlayer(scanner)
	p2 := scanPlayer(scanner)
	return game{player1: p1, player2: p2}
}

func scanPlayer(scanner *bufio.Scanner) player {
	p := player{}
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "Player") {
			continue
		}
		if len(text) == 0 {
			break
		}
		p = append(p, misc.MustAtoi(text))
	}
	return p
}
