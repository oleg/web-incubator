package main

import (
	"bufio"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
	"strings"
)

func main() {
	rules := ParsePassRules(misc.MustOpen("day2/input.txt"))

	println(rules.count(PassRule.valid1))
	println(rules.count(PassRule.valid2))
}

type PassRule struct {
	n1, n2   int
	letter   string
	password string
}
type PassRules []PassRule

func (p PassRule) valid1() bool {
	count := strings.Count(p.password, p.letter)
	return p.n1 <= count && count <= p.n2
}

func (p PassRule) valid2() bool {
	n1Match := string(p.password[p.n1-1]) == p.letter
	n2Match := string(p.password[p.n2-1]) == p.letter
	return n1Match != n2Match
}

func (rs PassRules) count(f func(PassRule) bool) int {
	count := 0
	for _, rule := range rs {
		if f(rule) {
			count++
		}
	}
	return count
}

func ParsePassRules(reader io.Reader) PassRules {
	rules := make([]PassRule, 0, 10)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		rules = append(rules, ParsePassRule(scanner.Text()))
	}
	return rules
}

//rule is in format of '1-3 a: abcde'
func ParsePassRule(s string) PassRule {
	numAndLetAndWord := strings.Split(s, " ")
	num := numAndLetAndWord[0]
	letter := numAndLetAndWord[1][:1]
	word := numAndLetAndWord[2]

	n1AndN2 := strings.Split(num, "-")
	n1 := misc.MustAtoi(n1AndN2[0])
	n2 := misc.MustAtoi(n1AndN2[1])

	return PassRule{n1: n1, n2: n2, letter: letter, password: word}
}
