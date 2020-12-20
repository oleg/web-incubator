package main

import (
	"bufio"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
	"regexp"
	"strings"
)

func main() {
	rs, _, nts := parseInput(misc.MustOpen("day16/input.txt"))
	println(scanningErrorRate(rs, nts))
}

type ticket []int
type rule [][2]int

func (r *rule) validForField(field int) bool {
	for _, v := range *r {
		if v[0] <= field && field <= v[1] {
			return true
		}
	}
	return false
}

func scanningErrorRate(rules map[string]rule, tickets []ticket) int {
	rate := 0
	for _, ticket := range tickets {
		rate += notValidForAnyField(ticket, rules)
	}
	return rate
}

func notValidForAnyField(ticket ticket, rules map[string]rule) int {
	for _, f := range ticket {
		valid := false
		for _, r := range rules {
			if r.validForField(f) {
				valid = true
				break
			}
		}
		if !valid {
			return f
		}
	}
	return 0
}

func parseInput(reader io.Reader) (map[string]rule, ticket, []ticket) {
	rules := make(map[string]rule)
	yticket := make(ticket, 0)
	ntickets := make([]ticket, 0)

	state := 1
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		if text == "your ticket:" {
			state = 2
			continue
		}
		if text == "nearby tickets:" {
			state = 3
			continue
		}
		if state == 1 {
			name, rng := parseRule(text)
			rules[name] = rng
		}
		if state == 2 {
			yticket = parseTicket(text)
		}
		if state == 3 {
			ntickets = append(ntickets, parseTicket(text))
		}
	}

	return rules, yticket, ntickets
}

func parseTicket(text string) ticket {
	var t ticket
	data := strings.Split(text, ",")
	for _, v := range data {
		t = append(t, misc.MustAtoi(v))
	}
	return t
}

var ruleRegexp = regexp.MustCompile(`^([\w ]+)+: (\d+)+-(\d+)+ or (\d+)+-(\d+)+$`)

func parseRule(text string) (string, rule) {
	data := ruleRegexp.FindStringSubmatch(text)
	return data[1], [][2]int{
		{misc.MustAtoi(data[2]), misc.MustAtoi(data[3])},
		{misc.MustAtoi(data[4]), misc.MustAtoi(data[5])},
	}
}
