package main

import (
	"fmt"
	"regexp"
	"strings"
)

type srule struct {
	text  string
	rules string
}

func day19() {
	groups := readGroupedLines("input/19.txt")
	rules := groups[0]
	toCheck := groups[1]

	//fmt.Printf("%v\n%v\n", rules, toCheck)

	rx := parseRules(rules)
	println(rx)

	reg := regexp.MustCompile(rx)
	count := 0
	for _, t := range toCheck {
		if reg.MatchString(t) {
			println(t)
			count++
		}
	}
	println(count)
}

func parseRules(rules []string) string {
	rs := map[string]srule{}
	rx := regexp.MustCompile("(.*): (.*)")
	for _, r := range rules {
		matches := rx.FindStringSubmatch(r)
		var sr srule
		if matches[2][0] == '"' {
			sr.text = matches[2]
		} else {
			sr.rules = matches[2]
		}
		rs[matches[1]] = sr
	}

	// Fix rule 11
	max := 10
	var sb strings.Builder
	for i := 1; i <= max; i++ {
		for j := 0; j < i; j++ {
			sb.WriteString("42 ")
		}
		for j := 0; j < i; j++ {
			sb.WriteString("31")
			if j != i-1 {
				sb.WriteString(" ")
			}
		}
		if i != max {
			sb.WriteString(" | ")
		}
	}
	rs["11"] = srule{rules: sb.String()}

	return "^" + generateRegex(rs, "0") + "$"
}

func generateRegex(rules map[string]srule, state string) string {
	rule := rules[state]
	if rule.text != "" {
		return string(rule.text[1])
	}

	if state == "8" {
		return fmt.Sprintf("(%v)+", generateRegex(rules, "42"))
	}

	var sb strings.Builder
	if strings.Contains(rule.rules, "|") {
		idx := strings.Index(rule.rules, " | ")
		leftRules := copyMap(rules)
		leftRules[state] = srule{rules: rule.rules[:idx]}
		lr := generateRegex(leftRules, state)

		rightRules := copyMap(rules)
		rightRules[state] = srule{rules: rule.rules[idx+3:]}
		rr := generateRegex(rightRules, state)
		return fmt.Sprintf("((%s)|(%s))", lr, rr)
	} else {
		for _, r := range strings.Split(rule.rules, " ") {
			r = strings.Trim(r, " ")
			sb.WriteString(generateRegex(rules, r))
		}
	}

	return sb.String()
}

func copyMap(rules map[string]srule) map[string]srule {
	m := make(map[string]srule)
	for k, v := range rules {
		m[k] = v
	}
	return m
}
