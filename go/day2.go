package main

import (
	"fmt"
	"strings"
)

type rule struct {
	min      int
	max      int
	char     byte
	password string
}

func day2() {
	lines := readLines("input/input-2.txt")

	var rules []rule
	for _, line := range lines {
		rules = append(rules, parseRule(line))
	}

	valid := 0
	for _, r := range rules {
		if !isValidComplexRule(r) {
			fmt.Printf("%v\n", r)
		} else {
			valid++
		}
	}
	println(valid)
}

func isValidRule(r rule) bool {
	num := strings.Count(r.password, string(r.char))
	return num >= r.min && num <= r.max
}

func isValidComplexRule(r rule) bool {
	c1 := r.password[r.min-1]
	c2 := r.password[r.max-1]

	if c1 == r.char && c2 != r.char {
		return true
	}

	if c1 != r.char && c2 == r.char {
		return true
	}

	return false
}

func parseRule(line string) rule {
	var r rule
	fmt.Sscanf(line, "%d-%d %c: %s", &r.min, &r.max, &r.char, &r.password)
	return r
}
