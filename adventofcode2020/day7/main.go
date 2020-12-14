package main

import (
	"bufio"
	"fmt"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
	"regexp"
	"strings"
)

func main() {
	reader := misc.MustOpen("day7/input.txt")
	bags := parseBags(reader)
	fmt.Println(countCanContain("shiny gold", bags))
}

type bag struct {
	name string
	bags []*bag
}

type bagsRepository map[string]*bag

func (bags bagsRepository) getOrCreate(name string) *bag {
	aBag, found := bags[name]
	if !found {
		aBag = &bag{name: name, bags: make([]*bag, 0)}
		bags[name] = aBag
	}
	return aBag
}

func (bags bagsRepository) newBag(rule string) *bag {
	topName := strings.Split(rule, " bags contain ")[0]
	innerBagNames := regexp.
		MustCompile(`\d (\w+ \w+)+`).
		FindAllStringSubmatch(rule, -1)

	topBag := bags.getOrCreate(topName)

	for _, name := range innerBagNames {
		aBag := bags.getOrCreate(name[1])
		topBag.bags = append(topBag.bags, aBag)
	}
	return topBag
}

func parseBags(reader io.Reader) []*bag {
	topLevel := make([]*bag, 0)
	bags := bagsRepository{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		topLevel = append(topLevel, bags.newBag(scanner.Text()))
	}
	return topLevel
}

func countCanContain(name string, bags []*bag) int {
	count := 0
	for _, b := range bags {
		if containAny(name, b) {
			count++
		}
	}
	return count - 1
}

func containAny(name string, bag *bag) bool {
	if bag.name == name {
		return true
	}
	for _, b := range bag.bags {
		if containAny(name, b) {
			return true
		}
	}
	return false
}
