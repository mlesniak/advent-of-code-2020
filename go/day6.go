package main

// Count number of yes votes for each character
type Group struct {
	Size   int
	Counts map[byte]int
}

func day6() {
	glines := readGroupedLines("input/6.txt")
	counts := parseGroup(glines)

	day6part1(counts)
	day6part2(counts)
}

func day6part2(groups []Group) {
	sum := 0
	for _, group := range groups {
		for _, v := range group.Counts {
			if v == group.Size {
				sum++
			}
		}
	}

	println(sum)
}

func day6part1(groups []Group) {
	sum := 0
	for _, group := range groups {
		sum += len(group.Counts)
	}
	println(sum)
}

func parseGroup(groupedLine [][]string) []Group {
	var groups []Group

	for _, group := range groupedLine {
		counts := make(map[byte]int)
		for _, member := range group {
			for i := range member {
				vote := member[i]
				counts[vote] = counts[vote] + 1
			}
		}

		groups = append(groups, Group{
			Size:   len(group),
			Counts: counts,
		})
	}

	return groups
}
