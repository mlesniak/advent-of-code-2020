package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	amount   int
	name     string
	children []Bag
}

func (b Bag) String() string {
	return fmt.Sprintf("%dx%s", b.amount, b.name)
}

func day7() {
	bagDefinitions := readLines("input/7.txt")
	bags := parseBagDefinitions(bagDefinitions)

	// Part1
	target := "shiny gold"
	count := 0

nextBag:
	for _, bag := range bags {
		// Direct?
		for _, child := range bag.children {
			if child.name == target {
				count++
				continue nextBag
			}
		}

		// Indirect over a child
		for _, child := range bag.children {
			for _, grandchild := range bags[child.name].children {
				if grandchild.name == target {
					count++
					continue nextBag
				}
			}
		}
	}

	println(count)
}

func parseBagDefinitions(lines []string) map[string]Bag {
	bags := make(map[string]Bag)

	regHasBag := regexp.MustCompile(`(.*) bags contain (.*)`)
	for _, line := range lines {
		matches := regHasBag.FindStringSubmatch(line)
		name := matches[1]
		rest := matches[2]
		children := parseChildren(rest)

		bags[name] = Bag{
			amount:   1,
			name:     name,
			children: children,
		}
	}

	return bags
}

func parseChildren(def string) []Bag {
	var bags []Bag
	if def == "no other bags." {
		return bags
	}

	bagDef := regexp.MustCompile(`(\d+) (.*) `)
	bs := strings.Split(def, ", ")
	for _, bd := range bs {
		matches := bagDef.FindStringSubmatch(bd)
		amount, err := strconv.Atoi(matches[1])
		name := matches[2]
		if err != nil {
			panic(err)
		}

		bags = append(bags, Bag{
			amount:   amount,
			name:     name,
			children: nil,
		})
	}

	return bags
}
