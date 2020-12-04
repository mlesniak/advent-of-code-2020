package main

import (
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	data map[string]string
}

func day4() {
	passports := readPassports("input/4.txt")

	count := 0
	for _, v := range passports {
		if isValid(v) && isValidComplex(v) {
			count++
		}
	}
	println(count)
}

func isValidComplex(p passport) bool {
	if !checkNumber(p, "byr", 1920, 2002) {
		return false
	}

	if !checkNumber(p, "iyr", 2010, 2020) {
		return false
	}

	if !checkNumber(p, "eyr", 2020, 2030) {
		return false
	}

	// Next time: more regular expression sub-group matching.
	hgt, ok := p.data["hgt"]
	if !ok {
		return false
	}
	if strings.HasSuffix(hgt, "cm") {
		s := hgt[:len(hgt)-2]
		nbyr, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if !(nbyr >= 150 && nbyr <= 193) {
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		s := hgt[:len(hgt)-2]
		nbyr, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if !(nbyr >= 59 && nbyr <= 76) {
			return false
		}
	} else {
		return false
	}

	if !checkRegex(p, "hcl", `#[a-f0-9]{6}`) {
		return false
	}

	if !checkSet(p, "ecl", []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}) {
		return false
	}

	if !checkRegex(p, "pid", `^\d{9}$`) {
		return false
	}

	return true
}

func checkSet(p passport, field string, valids []string) bool {
	byr, ok := p.data[field]
	if !ok {
		return false
	}

	return stringInSlice(byr, valids)
}

func stringInSlice(s string, slice []string) bool {
	for _, v := range slice {
		if s == v {
			return true
		}
	}
	return false
}

func checkNumber(p passport, field string, min, max int) bool {
	byr, ok := p.data[field]
	if !ok {
		return false
	}
	nbyr, err := strconv.Atoi(byr)
	if err != nil {
		return false
	}
	if !(nbyr >= min && nbyr <= max) {
		return false
	}

	return true
}

func checkRegex(p passport, field string, rx string) bool {
	byr, ok := p.data[field]
	if !ok {
		return false
	}

	r := regexp.MustCompile(rx)
	return r.MatchString(byr)
}

func isValid(p passport) bool {
	if len(p.data) == 8 {
		return true
	}

	_, found := p.data["cid"]
	if len(p.data) == 7 && !found {
		return true
	}

	return false
}

func readPassports(filename string) []passport {
	var passports []passport

	// Next time: simply split whole string by \n\n
	lines := readLines(filename)
	var pdata []string
	for _, line := range lines {
		if line == "" {
			jpd := strings.Join(pdata, " ")
			p := parsePassport(jpd)
			passports = append(passports, p)
			pdata = []string{}
		} else {
			pdata = append(pdata, line)
		}
	}

	return passports
}

func parsePassport(pdata string) passport {
	data := make(map[string]string)

	parts := strings.Split(pdata, " ")
	for _, v := range parts {
		pd := strings.Split(v, ":")
		data[pd[0]] = pd[1]
	}

	return passport{
		data: data,
	}
}
