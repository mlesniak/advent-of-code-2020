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

	//part1(bags)
	part2(bags)
}

func part2(bags map[string]Bag) {
	target := "shiny gold"
	count := countBag(bags, bags[target])
	// You don't need the shiny gold bag to carry itself.
	count = count - 1

	println(count)
}

func countBag(bags map[string]Bag, bag Bag) int {
	count := 0

	fmt.Printf("- %s\n", bag.name)
	// The bag itself
	count++

	// Count all children
	for _, child := range bag.children {
		fmt.Printf("\t%v\n", child)
		count = count + child.amount*countBag(bags, bags[child.name])
	}

	return count
}

func part1(bags map[string]Bag) {
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

		// Indirect over one of the children
		for _, child := range bag.children {
			if reachable(target, bags, bags[child.name].children) {
				count++
				continue nextBag
			}
		}
	}

	println(count)
}

func reachable(target string, bags map[string]Bag, nodes []Bag) bool {
	for _, node := range nodes {
		if node.name == target {
			return true
		}

		isReachableFor := reachable(target, bags, bags[node.name].children)
		if isReachableFor {
			return true
		}
	}

	return false
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
