package main

import (
	"fmt"
	"strings"
)

type passport struct {
	data map[string]string
}

func day4() {
	passports := readPassports("input/4.txt")

	count := 0
	for _, v := range passports {
		if isValid(v) {
			count++
		} else {
			fmt.Printf("%v\n", v)
		}
	}
	println(count)

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
