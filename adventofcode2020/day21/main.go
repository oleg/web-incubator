package main

import (
	"bufio"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"github.com/oleg/incubator/adventofcode2020/strcol"
	"io"
	"regexp"
	"sort"
	"strings"
)

func main() {
	r := parseInput(misc.MustOpen("day21/input.txt"))
	r.reduce()
	println(r.countMaybeNonAllergicIngredient())
	println(r.getCanonicalDangerousIngredientList())
}

type registry struct {
	a2i     map[string][][]string
	mapping map[string]string
	i       map[string]int
}

func (r *registry) intersects() map[string][]string {
	res := make(map[string][]string)
	for alg, ingsList := range r.a2i {
		res[alg] = intersection(ingsList)
	}
	return res
}

func (r *registry) update() {
	for k, v := range r.intersects() {
		if len(v) == 1 {
			r.mapping[k] = v[0]
		}
	}
}

func (r *registry) eliminate() {
	for k, v := range r.mapping {
		for lk, lv := range r.a2i {
			if k != lk {
				for i, l := range lv {
					lv[i] = strcol.Remove(l, v)
				}
			}
		}
	}
}

func (r *registry) reduce() {
	prevLen := -1
	for len(r.mapping) > prevLen {
		prevLen = len(r.mapping)
		r.update()
		r.eliminate()
	}
}
func (r *registry) countMaybeNonAllergicIngredient() int {
	allergicIngredients := strcol.Values(r.mapping)
	count := 0
	for k, v := range r.i {
		if !strcol.Contains(allergicIngredients, k) {
			count += v
		}
	}
	return count
}

func (r *registry) getCanonicalDangerousIngredientList() string {
	keys := strcol.Keys(r.mapping)
	sort.Strings(keys)
	var b strings.Builder
	b.WriteString(r.mapping[keys[0]])
	for _, k := range keys[1:] {
		b.WriteString(",")
		b.WriteString(r.mapping[k])
	}
	return b.String()
}

func intersection(lists [][]string) []string {
	ins := make([]string, 0)
	for _, item := range lists[0] {
		everywhere := true
		for _, otherList := range lists[1:] {
			if !strcol.Contains(otherList, item) {
				everywhere = false
				break
			}
		}
		if everywhere {
			ins = append(ins, item)
		}
	}
	return ins
}

var a2iRe = regexp.MustCompile(`([\s\w]+)+ \(contains ([\s\w,]+)+\)`)

func parseInput(reader io.Reader) registry {
	is := make(map[string]int)
	a2i := make(map[string][][]string)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		matches := a2iRe.FindAllStringSubmatch(scanner.Text(), -1)
		ings := strings.Split(matches[0][1], " ")
		for _, i := range ings {
			is[i]++
		}
		algns := strings.Split(matches[0][2], ", ")
		for _, algn := range algns {
			ingsList, found := a2i[algn]
			if !found {
				ingsList = make([][]string, 0)
			}
			a2i[algn] = append(ingsList, ings)
		}
	}
	return registry{a2i: a2i, i: is, mapping: make(map[string]string)}
}
