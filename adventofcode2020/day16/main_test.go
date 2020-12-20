package main

import (
	"strings"
	"testing"
)

var testData = strings.TrimPrefix(`
class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`, "\n")

func Test_day16_task1_parse(t *testing.T) {
	expectedRules := map[string]rule{
		"class": {{1, 3}, {5, 7}},
		"row":   {{6, 11}, {33, 44}},
		"seat":  {{13, 40}, {45, 50}},
	}
	expectedYourTicket := []int{7, 1, 14}
	expectedNearbyTickets := []ticket{
		{7, 3, 47},
		{40, 4, 50},
		{55, 2, 20},
		{38, 6, 12},
	}

	rules, yourTicket, nearbyTickets := parseInput(strings.NewReader(testData))

	if len(rules) != len(expectedRules) || rules["seat"][1][1] != expectedRules["seat"][1][1] {
		t.Errorf("Wrong rule %v", rules)
	}
	if len(yourTicket) != len(expectedYourTicket) || yourTicket[1] != expectedYourTicket[1] {
		t.Errorf("Wrong yourTicket %v", yourTicket)
	}
	if len(nearbyTickets) != len(expectedNearbyTickets) || nearbyTickets[1][1] != expectedNearbyTickets[1][1] {
		t.Errorf("Wrong nearbyTickets %v", nearbyTickets)
	}
}

func Test_day16_task1_count_invalid(t *testing.T) {
	rules, _, nearbyTickets := parseInput(strings.NewReader(testData))

	val := scanningErrorRate(rules, nearbyTickets)

	if val != 71 {
		t.Errorf("Wrong scanningErrorRate %v, expected %v", val, 71)
	}
}
