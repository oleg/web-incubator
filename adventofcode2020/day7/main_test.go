package main

import (
	"strings"
	"testing"
)

var testData = strings.TrimPrefix(`
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`, "\n")

func Test_day7_task1_parse_rule_multiple_inner(t *testing.T) {
	rule := `vibrant silver bags contain 4 dotted lavender bags, 3 wavy green bags, 1 striped yellow bag, 4 muted plum bags.`

	bag := bagsRepository{}.newBag(rule)

	if bag.name != "vibrant silver" {
		t.Errorf("Wrong name %s", bag.name)
	}
	if len(bag.bags) != 4 {
		t.Errorf("Wrong inner length %d", len(bag.bags))
	}
	if bag.bags[0].name != "dotted lavender" ||
		bag.bags[1].name != "wavy green" ||
		bag.bags[2].name != "striped yellow" ||
		bag.bags[3].name != "muted plum" {
		t.Errorf("Wrong inner name %v", bag.bags)
	}
}

func Test_day7_task1_parse_one_rule_no_inner(t *testing.T) {
	rule := `faded blue bags contain no other bags`

	bag := bagsRepository{}.newBag(rule)

	if bag.name != "faded blue" {
		t.Errorf("Wrong name %s", bag.name)
	}
	if len(bag.bags) != 0 {
		t.Errorf("Wrong contains number %d", len(bag.bags))
	}
}

func Test_day7_task1_parse_rule_reuse_bags(t *testing.T) {
	rule := "bright white bags contain 1 shiny gold bag."
	shinyGold := &bag{
		name: "shiny gold",
		bags: []*bag{{name: "inner1"}, {name: "inner2"}, {name: "inner3"}}}

	bags := map[string]*bag{"shiny gold": shinyGold}

	bag := bagsRepository(bags).newBag(rule)

	if len(bag.bags[0].bags) != 3 {
		t.Errorf("Bag was not reused %v", bag.bags[0])
	}
}

func Test_day7_task1_parse_rules(t *testing.T) {
	bags := parseBags(strings.NewReader(testData))
	if len(bags) != 9 {
		t.Errorf("Wrong number of top level bags %d", len(bags))
	}
	if bags[0].name != "light red" ||
		len(bags[0].bags) != 2 {
		t.Errorf("Wrong bag data %v", bags[0])
	}
	if bags[8].name != "dotted black" ||
		len(bags[8].bags) != 0 {
		t.Errorf("Wrong bag data %v", bags[8])
	}
}

func Test_day7_task1_count(t *testing.T) {
	bags := parseBags(strings.NewReader(testData))

	count := countCanContain("shiny gold", bags)

	if count != 4 {
		t.Errorf("Wrong count of bags that can contain 'shiny gold' %d", count)
	}
}
