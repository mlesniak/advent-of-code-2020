package main

import (
	"fmt"
	"strings"
)

type rule struct {
	min      int
	max      int
	char     rune
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
		if !isValidRule(r) {
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

func parseRule(line string) rule {
	var r rule
	fmt.Sscanf(line, "%d-%d %c: %s", &r.min, &r.max, &r.char, &r.password)
	return r
}
