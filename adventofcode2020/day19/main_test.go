package main

import (
	"fmt"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"strings"
	"testing"
)

func Test_day19_task1_constNode(t *testing.T) {
	cs := constNode{"a"}

	res := cs.data(repo{})

	if len(res) != 1 || res[0] != "a" {
		t.Errorf("Wrong value %v", res)
	}
}

func Test_day19_task1_seqNode_1(t *testing.T) {

	cs := seqNode{[]string{"1"}}

	res := cs.data(repo{
		"1": &constNode{"x"},
	})

	if len(res) != 1 || res[0] != "x" {
		t.Errorf("Wrong value %v", res)
	}
}

func Test_day19_task1_seqNode_2(t *testing.T) {
	r := repo{
		"1": &constNode{"a"},
		"2": &constNode{"b"},
		"3": &seqNode{[]string{"1", "2"}},
	}
	res := r["3"].data(r)
	res = r["3"].data(r)

	if len(res) != 1 || res[0] != "ab" {
		t.Errorf("Wrong value %v", res)
	}
}

func Test_day19_task1_seqNode_3(t *testing.T) {
	r := repo{
		"1": &constNode{"a"},
		"2": &constNode{"x"},
		"3": &constNode{"y"},
		"4": &altNode{[]string{"2", "3"}, []string{"3", "2"}},
		"5": &constNode{"b"},
		"6": &seqNode{[]string{"1", "4", "5"}},
	}
	res := r["6"].data(r)

	if len(res) != 2 || res[0] != "axyb" || res[1] != "ayxb" {
		t.Errorf("Wrong value %v", res)
	}
}

func Test_day19_task1_altNode(t *testing.T) {

	r := repo{
		"1": &constNode{"a"},
		"2": &constNode{"b"},
		"3": &altNode{idsA: []string{"1"}, idsB: []string{"2"}},
	}
	res := r["3"].data(r)

	if len(res) != 2 || res[0] != "a" || res[1] != "b" {
		t.Errorf("Wrong value %v", res)
	}
}

func Test_day19_task1_altNode_2(t *testing.T) {
	r := repo{
		"1": &constNode{"a"},
		"2": &constNode{"b"},
		"3": &constNode{"x"},
		"4": &constNode{"y"},
		"5": &seqNode{[]string{"1", "2"}},
		"6": &seqNode{[]string{"3", "4"}},
		"7": &altNode{idsA: []string{"5"}, idsB: []string{"6"}},
	}
	res := r["7"].data(r)

	if len(res) != 2 || res[0] != "ab" || res[1] != "xy" {
		t.Errorf("Wrong value %v", res)
	}
}

func Test_day19_task1_parse_rules(t *testing.T) {
	testData := misc.TrimNewLine(`
0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"`)

	r, _ := parseRulesAndData(strings.NewReader(testData))
	if len(r) != 6 {
		t.Errorf("Wrong size %v", r)
	}
	fmt.Printf("%v\n", r["1"].data(r))
}

func Test_day19_task1_parse_data(t *testing.T) {
	testData := misc.TrimNewLine(`
0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`)

	_, d := parseRulesAndData(strings.NewReader(testData))
	if len(d) != 5 || d[2] != "abbbab" {
		t.Errorf("Wrong size %v", d)
	}
}

func Test_day19_task1_count_matching(t *testing.T) {
	testData := misc.TrimNewLine(`
0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`)

	n := countMatching(parseRulesAndData(strings.NewReader(testData)))
	if n != 2 {
		t.Errorf("Wrong count %v", n)
	}
}
