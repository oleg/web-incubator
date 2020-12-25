package main

import (
	"bufio"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
	"regexp"
	"strings"
)

func main() {
	println(countMatching(parseRulesAndData(misc.MustOpen("day19/input.txt"))))
}

type repo map[string]node

type node interface {
	data(r repo) []string
}

type constNode struct {
	str string
}

func (c *constNode) data(r repo) []string {
	return []string{c.str}
}

type seqNode struct {
	ids []string
}

func (c *seqNode) data(r repo) []string {
	if len(c.ids) == 0 {
		return nil
	}
	parts := make([][]string, len(c.ids))
	for i, id := range c.ids {
		parts[i] = r[id].data(r)
	}
	return join(parts[0], parts[1:])
}

func join(head []string, rest [][]string) []string {
	if len(rest) == 0 {
		return head
	}
	res := make([]string, 0)
	for _, suffix := range head {
		for _, prefix := range rest[0] {
			res = append(res, suffix+prefix)
		}
	}
	return join(res, rest[1:])
}

type altNode struct {
	idsA, idsB []string
}

func (n *altNode) data(r repo) []string {
	res := (&seqNode{n.idsA}).data(r)
	res = append(res, (&seqNode{n.idsB}).data(r)...)
	return res
}

var altStrRe = regexp.MustCompile(`^(\d+)+:([\s\d]+)+\|([\s\d]+)+$`)
var seqStrRe = regexp.MustCompile(`^(\d+)+:([\s\d]+)+$`)
var constStrRe = regexp.MustCompile(`^(\d+)+: "(\w+)"$`)

func parseRulesAndData(reader io.Reader) (repo, []string) {
	rules := repo{}
	data := make([]string, 0)
	scanner := bufio.NewScanner(reader)
	scanRules := true
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			scanRules = false
			continue
		}
		if scanRules {
			if cnst := constStrRe.FindAllStringSubmatch(text, -1); len(cnst) > 0 {
				rules[cnst[0][1]] = &constNode{
					cnst[0][2],
				}
			}
			if seq := seqStrRe.FindAllStringSubmatch(text, -1); len(seq) > 0 {
				rules[seq[0][1]] = &seqNode{
					strings.Split(strings.Trim(seq[0][2], " "), " "),
				}
			}
			if alt := altStrRe.FindAllStringSubmatch(text, -1); len(alt) > 0 {
				rules[alt[0][1]] = &altNode{
					strings.Split(strings.Trim(alt[0][2], " "), " "),
					strings.Split(strings.Trim(alt[0][3], " "), " "),
				}
			}
		} else {
			data = append(data, text)
		}
	}
	return rules, data
}

func countMatching(r repo, data []string) int {
	vl := r["0"].data(r)
	vm := make(map[string]struct{}, len(vl))
	for _, v := range vl {
		vm[v] = struct{}{}
	}
	count := 0
	for _, v := range data {
		if _, found := vm[v]; found {
			count++
		}
	}
	return count
}
