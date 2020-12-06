package main

import (
	"bufio"
	"github.com/oleg/incubator/adventofcode2020/misc"
	"io"
	"strconv"
	"strings"
)

func main() {
	passports := ParsePassports(misc.MustOpen("day4/input.txt"))
	count1 := 0
	count2 := 0
	for _, p := range passports {
		if p.valid() {
			count1++
		}
		if p.strictValid() {
			count2++
		}
	}
	println(count1)
	println(count2)
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
		res := strings.Split(text, " ")
		for _, kv := range res {
			a := strings.Split(kv, ":")
			if len(a) == 2 {
				pass[a[0]] = a[1]
			}
		}
	}
	passports = append(passports, pass)
	return passports
}

type Passport map[string]string

func (p Passport) valid() bool {
	_, byrFound := p["byr"]
	_, iyrFound := p["iyr"]
	_, eyrFound := p["eyr"]
	_, hgtFound := p["hgt"]
	_, hclFound := p["hcl"]
	_, eclFound := p["ecl"]
	_, pidFound := p["pid"]
	//_, cidFound := p["cid"]
	return byrFound && iyrFound && eyrFound && hgtFound && hclFound && eclFound && pidFound
}

func (p Passport) strictValid() bool {
	return p.byrValid() && p.iyrValid() && p.eyrValid() && p.hgtValid() && p.hclValid() && p.eclValid() && p.pidValid()
}

func (p Passport) byrValid() bool {
	if val, found := p["byr"]; found {
		if len(val) == 4 {
			if num, err := strconv.Atoi(val); err == nil {
				if 1920 <= num && num <= 2002 {
					return true
				}
			}

		}
	}
	return false
}

func (p Passport) iyrValid() bool {
	if val, found := p["iyr"]; found {
		if len(val) == 4 {
			if num, err := strconv.Atoi(val); err == nil {
				if 2010 <= num && num <= 2020 {
					return true
				}
			}

		}
	}
	return false
}

func (p Passport) eyrValid() bool {
	if val, found := p["eyr"]; found {
		if len(val) == 4 {
			if num, err := strconv.Atoi(val); err == nil {
				if 2020 <= num && num <= 2030 {
					return true
				}
			}

		}
	}
	return false
}

func (p Passport) hgtValid() bool {
	if val, found := p["hgt"]; found {
		if strings.HasSuffix(val, "cm") {
			if num, err := strconv.Atoi(val[:len(val)-2]); err == nil {
				if 150 <= num && num <= 193 {
					return true
				}
			}
		}
		if strings.HasSuffix(val, "in") {
			if num, err := strconv.Atoi(val[:len(val)-2]); err == nil {
				if 59 <= num && num <= 76 {
					return true
				}
			}
		}
	}
	return false
}

func (p Passport) hclValid() bool {
	if val, found := p["hcl"]; found {
		if strings.HasPrefix(val, "#") {
			if _, err := strconv.ParseInt(val[1:], 16, 0); err == nil {
				return true
			}
		}
	}
	return false
}

func (p Passport) eclValid() bool {
	if val, found := p["ecl"]; found {
		for _, v := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
			if val == v {
				return true
			}
		}

	}
	return false
}

func (p Passport) pidValid() bool {
	if val, found := p["pid"]; found {
		if len(val) == 9 {
			if _, err := strconv.Atoi(val); err == nil {
				return true
			}
		}
	}
	return false
}
