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
	rules, myTicket, otherTicket := readTickets("input/16.txt")
	fmt.Printf("%v\n", rules)
	fmt.Printf("%v\n", myTicket)
	fmt.Printf("%v\n", otherTicket)

	validTickets := computeValidTickets(otherTicket, rules)

	fmt.Printf("%v\n", validTickets)
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
