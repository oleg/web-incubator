package main

import (
	"bufio"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
	"strings"
)

func main() {
	passports := ParsePassports(misc.MustOpen("day4/input.txt"))
	count := 0
	for _, p := range passports {
		if p.valid() {
			count++
		}
	}
	println(count)
}

func ParsePassports(reader io.Reader) []Passport {
	scanner := bufio.NewScanner(reader)
	passports := make([]Passport, 0)
	pass := Passport{}
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			passports = append(passports, pass)
			pass = Passport{}
		}
		if strings.Contains(text, "byr:") {
			pass.byr = true
		}
		if strings.Contains(text, "iyr:") {
			pass.iyr = true
		}
		if strings.Contains(text, "eyr:") {
			pass.eyr = true
		}
		if strings.Contains(text, "hgt:") {
			pass.hgt = true
		}
		if strings.Contains(text, "hcl:") {
			pass.hcl = true
		}
		if strings.Contains(text, "ecl:") {
			pass.ecl = true
		}
		if strings.Contains(text, "pid:") {
			pass.pid = true
		}
		if strings.Contains(text, "cid:") {
			pass.cid = true
		}
	}
	passports = append(passports, pass)
	return passports
}

type Passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid bool
}

func (p *Passport) valid() bool {
	return p.byr && p.iyr && p.eyr && p.hgt && p.hcl && p.ecl && p.pid
}
