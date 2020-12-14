package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func day14() {
	commands := readLines("input/14.txt")
	memory := make(map[int]int)

	maskRegex := regexp.MustCompile(`mask = (.*)`)
	memRegex := regexp.MustCompile(`mem\[(\d)+\] = (\d+)`)
	for _, command := range commands {
		m1 := maskRegex.FindStringSubmatch(command)
		if len(m1) > 1 {
			mask := m1[1]
			fmt.Printf("%v\n", mask)
			continue
		}

		m2 := memRegex.FindStringSubmatch(command)
		if len(m2) > 1 {
			dst, _ := strconv.Atoi(m2[1])
			val, _ := strconv.Atoi(m2[2])
			fmt.Printf("%v <- %v\n", dst, val)
			continue
		}

		panic("Unknown command:" + command)
	}

	// Result
	sum := 0
	for _, v := range memory {
		sum += v
	}
	println(sum)
}
