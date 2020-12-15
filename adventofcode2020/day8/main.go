package main

import (
	"bufio"
	"errors"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
	"strings"
)

func main() {
	c1 := parseComputer(misc.MustOpen("day8/input.txt"))
	c1.execute()
	println(c1.accumulator)

	c2 := parseComputer(misc.MustOpen("day8/input.txt"))
	c2.healingExecute()
	println(c2.accumulator)
}

type computer struct {
	instructions []instruction
	visited      []bool
	accumulator  int
}

func (c *computer) execute() error {
	for pos := 0; pos < len(c.instructions); pos++ {
		if c.visited[pos] {
			return errors.New("infinite loop")
		}
		c.visited[pos] = true
		i := c.instructions[pos]
		switch i.operation {
		case "nop":
			continue
		case "acc":
			c.accumulator += i.arg
		case "jmp":
			pos += i.arg - 1
		}
	}
	return nil
}

func (c *computer) healingExecute() {
	swap := -1
	for err := errors.New("start"); err != nil; err = c.execute() {
		c.accumulator = 0
		c.visited = make([]bool, len(c.instructions))
		swap = swapAfter(c, swap)
	}
}

func swapAfter(c *computer, swapAfter int) int {
	if swapAfter != -1 {
		swap(&c.instructions[swapAfter])
	}
	for i := swapAfter + 1; i < len(c.instructions); i++ {
		if swap(&c.instructions[i]) {
			return i
		}
	}
	return -1
}

func swap(i *instruction) bool {
	if i.operation == "nop" {
		i.operation = "jmp"
		return true
	}
	if i.operation == "jmp" {
		i.operation = "nop"
		return true
	}
	return false
}

type instruction struct {
	operation string
	arg       int
}

func parseComputer(reader io.Reader) computer {
	scanner := bufio.NewScanner(reader)
	instructions := make([]instruction, 0)
	for scanner.Scan() {
		instructions = append(instructions, parseInstruction(scanner.Text()))
	}
	visited := make([]bool, len(instructions))
	return computer{instructions: instructions, visited: visited}
}

func parseInstruction(text string) instruction {
	operationAndArgument := strings.Split(text, " ")
	operation := operationAndArgument[0]
	argument := misc.MustAtoi(operationAndArgument[1])

	return instruction{operation: operation, arg: argument}
}
