package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_day24_center(t *testing.T) {
	r := newRepo()

	c := r.center()
	if c.key() != "0,0" {
		t.Errorf("Wrong key tile %v", c.key())
	}
}

func Test_day24_extend(t *testing.T) {
	r := newRepo()

	c := r.center()

	t1 := r.extend(c, "nw")
	if t1 == nil {
		t.Errorf("tile must not be nil")
	}
	if t1.key() != "1,-1" {
		t.Errorf("wrong k1.key(): %v", t1.key())
	}

	if c.nw == nil {
		t.Errorf("wrong c.nw value %+v", c.nw)
	}
	if c.nw.key() != "1,-1" {
		t.Errorf("wrong c.nw.key(): %v", c.nw.key())
	}
	if c.nw.se != c {
		t.Errorf("wrong c.nw.se value %+v", c.nw.se)
	}
}

func Test_day24_splitDirs(t *testing.T) {

	dirs := splitDirs("eseswwnwne")

	if !reflect.DeepEqual(dirs, []string{"e", "se", "sw", "w", "nw", "ne"}) {
		t.Errorf("Wrongly dirs dirs %v", dirs)
	}
}

func Test_day24_walk1(t *testing.T) {
	r := newRepo()
	t1 := r.walk("esew")
	if t1.key() != "-1,1" {
		t.Errorf("Found wrong tile %v", t)
	}
}

func Test_day24_walk2(t *testing.T) {
	r := newRepo()
	t1 := r.walk("nwwswee")
	if t1.key() != "0,0" {
		t.Errorf("Found wrong tile %v", t)
	}
}

func Test_day24_larger_example(t *testing.T) {
	testdata := `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`
	count := flipAndCountBlack(strings.NewReader(testdata))

	if count != 10 {
		t.Errorf("Wrong count of black tiles %v, expected %v", count, 10)
	}
}
