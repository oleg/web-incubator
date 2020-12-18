package main

import (
	"bufio"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
)

func main() {
	reader := bufio.NewReader(misc.MustOpen("day12/input.txt"))
	instructions := parseInstructions(reader)
	s1 := ship{}
	s1.doAll(instructions)
	println(s1.east, s1.north)

	s2 := ship{}
	s2.doAll2(instructions)
	println(s2.east, s2.north)
}

type instruction struct {
	action   string
	argument int
}

func parseInstructions(reader io.Reader) []instruction {
	instructions := make([]instruction, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		instructions = append(instructions, parseInstruction(scanner.Text()))
	}
	return instructions
}

func parseInstruction(text string) instruction {
	return instruction{action: text[:1], argument: misc.MustAtoi(text[1:])}
}

type location struct {
	east, north int
}

type ship struct {
	east, north, direction int
	waypoint               location
}

func (s *ship) doAll(is []instruction) {
	for _, i := range is {
		s.do(i)
	}
}

func (s *ship) do(i instruction) {
	switch i.action {
	case "N":
		s.north += i.argument
	case "S":
		s.north -= i.argument
	case "E":
		s.east += i.argument
	case "W":
		s.east -= i.argument
	case "L":
		s.direction = normalizeDirection(s.direction + (360 - i.argument))
	case "R":
		s.direction = normalizeDirection(s.direction + i.argument)
	case "F":
		switch s.direction {
		case 0:
			s.east += i.argument
		case 90:
			s.north -= i.argument
		case 180:
			s.east -= i.argument
		case 270:
			s.north += i.argument
		}
	}
}

func (s *ship) doAll2(is []instruction) {
	s.waypoint = location{east: 10, north: 1}
	for _, i := range is {
		s.do2(i)
	}
}

func (s *ship) do2(i instruction) {
	switch i.action {
	case "N":
		s.waypoint.north += i.argument
	case "S":
		s.waypoint.north -= i.argument
	case "E":
		s.waypoint.east += i.argument
	case "W":
		s.waypoint.east -= i.argument
	case "L":
		rotateRight(&s.waypoint, 360-i.argument)
	case "R":
		rotateRight(&s.waypoint, i.argument)
	case "F":
		s.east += s.waypoint.east * i.argument
		s.north += s.waypoint.north * i.argument
	}
}

func rotateRight(w *location, argument int) {
	for i := 0; i < argument/90; i++ {
		w.east, w.north = w.north, -w.east
	}
}

func normalizeDirection(d int) int {
	if d >= 360 {
		return d - 360
	}
	return d
}
