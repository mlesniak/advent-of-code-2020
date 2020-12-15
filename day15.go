package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day15() {
	lines := readLines("input/15.txt")
	nums := strings.Split(lines[0], ",")

	var lastNumberSpoken int
	lastSeen := make(map[int]int)
	var numbers []int
	for i, num := range nums {
		n, _ := strconv.Atoi(num)
		lastNumberSpoken = n
		if i != len(numbers)-1 {
			numbers = append(numbers, n)
			lastSeen[n] = i + 1
		}
	}

	turn := len(nums) + 1
	for {
		//fmt.Printf("\nturn=%d\n", turn)
		//fmt.Printf("lastNumberSpoken=%d\n", lastNumberSpoken)
		//fmt.Printf("memory=%v\n", lastSeen)

		// Check if spoken before:
		spokenBeforeAt := -1
		spokenBeforeAt, found := lastSeen[lastNumberSpoken]
		if !found || turn == len(numbers)+1 {
			spokenBeforeAt = -1
		}
		//fmt.Printf("spokenBeforeAt=%d\n", spokenBeforeAt)
		lastSeen[lastNumberSpoken] = turn - 1

		if spokenBeforeAt == -1 {
			//fmt.Printf("Never seen -> Adding 0\n")
			lastNumberSpoken = 0
		} else {
			x := turn - 1 - spokenBeforeAt
			//fmt.Printf("Seen before-> Adding %d\n", x)
			lastNumberSpoken = x
		}

		turn++
		if turn > 30000000 {
			fmt.Printf("\n\n-> %d\n", lastNumberSpoken)
			break
		}
	}
}
