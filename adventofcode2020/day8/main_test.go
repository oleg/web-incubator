package main

import (
	"strings"
	"testing"
)

var testData = strings.TrimPrefix(`
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`, "\n")

func Test_day8_task1_parse_instructions(t *testing.T) {
	instructions := parseComputer(strings.NewReader(testData)).instructions

	instLen := len(instructions)
	if instLen != 9 {
		t.Errorf("Wrong number of instructions read %d", instLen)
	}

	inst5 := instructions[5]
	if inst5.operation != "acc" || inst5.arg != -99 {
		t.Errorf("Wrong instruction %v", inst5)
	}
}

func Test_day8_execute_instructions(t *testing.T) {
	tests := []struct {
		program     string
		expectedAcc int
	}{
		{"nop +0", 0},
		{"nop +1", 0},
		{"acc +1", 1},
		{"acc -2", -2},
		{"acc +1\nnop +1\nacc +1", 2},
		{"jmp +1\nacc +1\nacc +1", 2},
		{"jmp +2\nacc +1\nacc +1", 1},
		{"jmp +2\njmp +2\njmp -1\nacc +5", 5},
	}
	for _, test := range tests {
		t.Run(test.program, func(t *testing.T) {
			comp := parseComputer(strings.NewReader(test.program))

			comp.execute()

			if comp.accumulator != test.expectedAcc {
				t.Errorf("Wrong value in acc %d, expected %d", comp.accumulator, test.expectedAcc)
			}
		})
	}
}

func Test_day8_task1_detect_loop(t *testing.T) {
	cmp := parseComputer(strings.NewReader(testData))

	cmp.execute()

	if cmp.accumulator != 5 {
		t.Errorf("Wrong value in accumulator %d", cmp.accumulator)
	}
}

func Test_day8_task2_self_healing_execute(t *testing.T) {
	cmp := parseComputer(strings.NewReader(testData))

	cmp.healingExecute()

	if cmp.accumulator != 8 {
		t.Errorf("Wrong value in accumulator %d", cmp.accumulator)
	}
}
