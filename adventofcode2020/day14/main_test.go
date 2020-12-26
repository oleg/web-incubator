package main

import (
	"strings"
	"testing"
)

func Test_day_14_task1_getMaskStr(t *testing.T) {
	m, present := getMaskStr("mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")

	if present != true || m != "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X" {
		t.Errorf("Wrong mask present: %v, str:%v", present, m)
	}
}

func Test_day14_task1_getMemStr(t *testing.T) {
	pos, num := getMem("mem[7] = 101")

	if pos != 7 || num != 101 {
		t.Errorf("wrong position %v or number %v", pos, num)
	}
}

func Test_day14_task1_newMask(t *testing.T) {
	m := newMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")

	if m.String() != "111111111111111111111111111111111101,000000000000000000000000000001000000" {
		t.Errorf("Wrong mask value %v", m)
	}
}

func Test_day14_task2(t *testing.T) {
	d := newDevice()
	d.setMask(newMask("000000000000000000000000000000X1001X"))
	d.setMem2(42, 100)

	sum := d.sum()

	if sum != 400 {
		t.Errorf("Wrong sum %v", sum)
	}
}

func Test_day14_task1_execute_instructions(t *testing.T) {
	instructions := strings.TrimPrefix(`
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`, "\n")

	d := newDevice()

	processInstructions(strings.NewReader(instructions), d.setMask, d.setMem)

	if d.mask.String() != newMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X").String() {
		t.Errorf("Wrong mask %v", d.mask)
	}
	if d.memory[7] != 101 || d.memory[8] != 64 {
		t.Errorf("Wrong device state %v", d.memory)
	}
	if d.sum() != 165 {
		t.Errorf("Wrong device memory sum %v", d.sum())
	}
}
