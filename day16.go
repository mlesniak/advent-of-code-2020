package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ticketRule struct {
	name string
	min  int
	max  int
}

type ticket struct {
	numbers []int
}

func day16() {
	rules, _, otherTickets := readTickets("input/16.txt")
	//fmt.Printf("%v\n", rules)
	//fmt.Printf("%v\n", myTicket)
	//fmt.Printf("%v\n", otherTickets)

	validTickets := computeValidTickets(otherTickets, rules)

	// For each ticketRule and position, check if possible at all.
	possRules := make(map[string]map[int]bool)
	for _, rule := range rules {
		possRules[rule[0].name] = make(map[int]bool)
	nextPosition:
		for pos := 0; pos < len(validTickets[0].numbers); pos++ {
			for _, ot := range validTickets {
				valPos := ot.numbers[pos]
				if !validTicketRule(valPos, rule) {
					continue nextPosition
				}
			}

			possRules[rule[0].name][pos] = true
		}

	}

	fmt.Printf("%v\n", possRules)
}

func validTicketRule(val int, rules []ticketRule) bool {
	foundValidRule := false
	for _, rule := range rules {
		if val >= rule.min && val <= rule.max {
			foundValidRule = true
			break
		}
	}

	return foundValidRule
}

func computeValidTickets(otherTicket []ticket, rules map[string][]ticketRule) []ticket {
	var validTickets []ticket
nextTicket:
	for _, t := range otherTicket {
		for _, tv := range t.numbers {
			validNumber := false
			// Check if this value is in any rule
		ruleCheck:
			for _, ticketRules := range rules {
				for _, r := range ticketRules {
					if tv <= r.max && tv >= r.min {
						validNumber = true
						break ruleCheck
					}
				}
			}
			if !validNumber {
				println(tv)
				continue nextTicket
			}
		}

		validTickets = append(validTickets, t)
	}
	return validTickets
}

func readTickets(filename string) (map[string][]ticketRule, ticket, []ticket) {
	lines := readGroupedLines(filename)

	rules := make(map[string][]ticketRule)
	for _, ruleLine := range lines[0] {
		parts := strings.Split(ruleLine, ":")
		name := parts[0]
		var rs []ticketRule
		for _, rl := range strings.Split(parts[1], " or ") {
			rl = strings.Trim(rl, " ")
			var r ticketRule
			r.name = name
			_, err := fmt.Sscanf(rl, "%d-%d", &r.min, &r.max)
			if err != nil {
				panic(err)
			}

			rs = append(rs, r)
		}
		rules[name] = rs
	}

	myTicket := parseTicket(lines[1][1])

	var tickets []ticket
	for i, t := range lines[2] {
		if i == 0 {
			continue
		}
		tickets = append(tickets, parseTicket(t))
	}

	return rules, myTicket, tickets
}

func parseTicket(line string) ticket {
	var ticket ticket
	for _, num := range strings.Split(line, ",") {
		n, _ := strconv.Atoi(num)
		ticket.numbers = append(ticket.numbers, n)
	}
	return ticket
}
