package main

import (
	"strings"
	"testing"
)

var testData = strings.TrimPrefix(`
abc

a
b
c

ab
ac

a
a
a
a

b`, "\n")

func Test_day6_task1_parse_groups(t *testing.T) {
	reader := strings.NewReader(testData)

	groups := parseGroups(reader)

	if len(groups) != 5 {
		t.Errorf("Wrong number of groups %d expected 5", len(groups))
	}

	if len(groups[0].lines) != 1 {
		t.Errorf("Wrong number of lines %d expected 1", len(groups[0].lines))
	}
	if len(groups[1].lines) != 3 {
		t.Errorf("Wrong number of lines %d expected 3", len(groups[1].lines))
	}
	if len(groups[2].lines) != 2 {
		t.Errorf("Wrong number of lines %d expected 2", len(groups[2].lines))
	}
	if len(groups[3].lines) != 4 {
		t.Errorf("Wrong number of lines %d expected 4", len(groups[3].lines))
	}
	if len(groups[4].lines) != 1 {
		t.Errorf("Wrong number of lines %d expected 5", len(groups[1].lines))
	}

	if groups[2].lines[0] != 3 {
		t.Errorf("Wrong contet of the line g[2].l[0] %d", groups[2].lines[0])
	}
	if groups[2].lines[1] != 5 {
		t.Errorf("Wrong contet of the line g[2].l[1] %d", groups[2].lines[1])
	}
}

func Test_day6_count_unique_answers(t *testing.T) {
	groups := parseGroups(strings.NewReader(testData))

	ua0 := countUniqueAnswers(groups[0])
	ua3 := countUniqueAnswers(groups[3])

	if ua0 != 3 {
		t.Errorf("Wrong number of unique answers %d expected 3", ua0)
	}
	if ua3 != 1 {
		t.Errorf("Wrong number of unique answers %d expected 1", ua3)
	}
}

func Test_day6_count_repeated_answers(t *testing.T) {
	groups := parseGroups(strings.NewReader(testData))

	ra0 := countRepeatedAnswers(groups[0])
	ra1 := countRepeatedAnswers(groups[1])
	ra2 := countRepeatedAnswers(groups[2])

	if ra0 != 3 {
		t.Errorf("Wrong number of repeated answers %d expected 3", ra0)
	}
	if ra1 != 0 {
		t.Errorf("Wrong number of repeated answers %d expected 0", ra1)
	}
	if ra2 != 1 {
		t.Errorf("Wrong number of repeated answers %d expected 1", ra2)
	}

}
