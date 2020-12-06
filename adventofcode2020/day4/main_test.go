package main

import (
	"strings"
	"testing"
)

var testPassportData = strings.Trim(`
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`, "\n")

func Test_day4_task1_parse_passports(t *testing.T) {
	passports := ParsePassports(strings.NewReader(testPassportData))

	passCount := len(passports)
	if passCount != 4 {
		t.Errorf("Wrong lengs of passports %d", passCount)
	}
}

func Test_day4_task1_valid(t *testing.T) {
	passports := ParsePassports(strings.NewReader(testPassportData))

	if !passports[0].valid() {
		t.Errorf("First passport must be valid")
	}
	if passports[1].valid() {
		t.Errorf("Second passport must be invalid")
	}
}
