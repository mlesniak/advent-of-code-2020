package main

// Count number of yes votes for each character
type Group = map[byte]int

func day6() {
	glines := readGroupedLines("input/6.txt")
	groups := parseGroup(glines)

	sum := 0
	for _, group := range groups {
		sum += len(group)
	}
	println(sum)
}

func parseGroup(groupedLine [][]string) []Group {
	var groups []Group

	for _, group := range groupedLine {
		g := make(Group)
		for _, member := range group {
			for i := range member {
				vote := member[i]
				g[vote] = g[vote] + 1
			}
		}

		groups = append(groups, g)
	}

	return groups
}
