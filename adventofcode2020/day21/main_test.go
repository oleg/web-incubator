package main

import (
	"fmt"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"github.com/oleg/incubator/adventofcode2020/strcol"
	"strings"
	"testing"
)

var testData = misc.TrimNewLine(`
mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`)

func Test_day21_task1_parse(t *testing.T) {
	r := parseInput(strings.NewReader(testData))

	if len(r.a2i) != 3 {
		t.Errorf("Wrong number of algns %v expected 3", len(r.a2i))
	}

	if len(r.a2i["fish"]) != 2 ||
		len(r.a2i["soy"]) != 1 {
		t.Errorf("Wrong a2i content %v", r.a2i)
	}
}

func Test_day21_task1_intersects(t *testing.T) {
	r := parseInput(strings.NewReader(testData))

	var ins = r.intersects()

	if !strcol.Eq(ins["fish"], []string{"mxmxvkd", "sqjhc"}) {
		t.Errorf("Wrong intersection content %v", ins["fish"])
	}
	if !strcol.Eq(ins["dairy"], []string{"mxmxvkd"}) {
		t.Errorf("Wrong intersection content %v", ins["dairy"])
	}
	if !strcol.Eq(ins["soy"], []string{"sqjhc", "fvjkl"}) {
		t.Errorf("Wrong intersection content %v", ins["soy"])
	}
}

func Test_day21_task1_update(t *testing.T) {
	r := parseInput(strings.NewReader(testData))

	r.update()

	if r.mapping["dairy"] != "mxmxvkd" {
		t.Errorf("Wrong mapping for dairy %v", r.mapping["dairy"])
	}
}

func Test_day21_task1_eliminate(t *testing.T) {
	r := parseInput(strings.NewReader(testData))

	r.update()
	r.eliminate()
	ins := r.intersects()

	fmt.Printf("%v\n", ins)

	if !strcol.Eq(ins["fish"], []string{"sqjhc"}) {
		t.Errorf("Wrong intersection content %v", ins["fish"])
	}
	if !strcol.Eq(ins["soy"], []string{"sqjhc", "fvjkl"}) {
		t.Errorf("Wrong intersection content %v", ins["soy"])
	}
}

func Test_day21_task1_reduce(t *testing.T) {
	r := parseInput(strings.NewReader(testData))

	r.reduce()
	count := r.countMaybeNonAllergicIngredient()

	if count != 5 {
		t.Errorf("Wrong count %v", count)
	}
}

func Test_day21_task2(t *testing.T) {
	r := parseInput(strings.NewReader(testData))

	r.reduce()
	res := r.getCanonicalDangerousIngredientList()

	if res != "mxmxvkd,sqjhc,fvjkl" {
		t.Errorf("Wrong dange list %v", res)
	}
}
