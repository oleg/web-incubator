package main

import (
	"github.com/oleg/incubator/adventofcode2020/misc"
	"strings"
	"testing"
)

var testData = misc.TrimNewLine(`
Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
`)

func Test_day22_task1_parse(t *testing.T) {
	g := parseGame(strings.NewReader(testData))

	if len(g.player1) != 5 || len(g.player2) != 5 {
		t.Errorf("Wrong number of cards %v %v", len(g.player1), len(g.player2))
	}
}

func Test_day22_task1_take(t *testing.T) {
	p := player{1, 2, 3}

	v := p.take()

	if v != 1 {
		t.Errorf("Wrong value %v", v)
	}
	if len(p) != 2 ||
		p[0] != 2 ||
		p[1] != 3 {
		t.Errorf("Wrong value %v", p)
	}
}

func Test_day22_task1_put(t *testing.T) {
	p := player{1, 2, 3}

	p.put(4, 5)

	if len(p) != 5 ||
		p[0] != 1 ||
		p[1] != 2 ||
		p[2] != 3 ||
		p[3] != 4 ||
		p[4] != 5 {
		t.Errorf("Wrong value %v", p)
	}
}

func Test_day22_task1_move(t *testing.T) {
	g := game{player1: player{5}, player2: player{1}}

	g.move()

	if len(g.player1) != 2 || len(g.player2) != 0 {
		t.Errorf("Wrong value %v", g)
	}
}

func Test_day22_task1_cycle(t *testing.T) {
	g := parseGame(strings.NewReader(testData))

	g.moveToEnd()

	if len(g.player1) != 0 || len(g.player2) != 10 {
		t.Errorf("Wrong value %v", g)
	}
}

func Test_day22_task1_count(t *testing.T) {
	g := parseGame(strings.NewReader(testData))

	g.moveToEnd()
	count := g.count()

	if count != 306 {
		t.Errorf("Wrong value %v", count)
	}
}
