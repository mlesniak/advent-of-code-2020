package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day15() {
	lines := readLines("input/15.txt")
	nums := strings.Split(lines[0], ",")

	var numbers []int
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		numbers = append(numbers, n)
	}

	turn := len(numbers) + 1
	for {
		fmt.Printf("\nturn=%d\n", turn)
		lastNumberSpoken := numbers[len(numbers)-1]
		fmt.Printf("lastNumberSpoken=%d\n", lastNumberSpoken)

		// Check if spoken before:
		spokenBeforeAt := -1
		for i := len(numbers) - 2; i >= 0; i-- {
			if numbers[i] == lastNumberSpoken {
				spokenBeforeAt = i
				break
			}
		}
		fmt.Printf("spokenBeforeAt=%d\n", spokenBeforeAt)

		if spokenBeforeAt == -1 {
			fmt.Printf("Never seen -> Adding 0\n")
			numbers = append(numbers, 0)
		} else {
			x := turn - 1 - spokenBeforeAt - 1
			fmt.Printf("Seen before-> Adding %d\n", x)
			numbers = append(numbers, x)
		}

		turn++
		if turn > 2020 {
			fmt.Printf("\n\n-> %d\n", numbers[len(numbers)-1])
			break
		}
	}
}
