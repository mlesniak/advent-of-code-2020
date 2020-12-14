package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	maskRegex = regexp.MustCompile(`mask = (.*)`)
	memRegex  = regexp.MustCompile(`mem\[(\d)+\] = (\d+)`)
)

func day14() {
	commands := readLines("input/14.txt")
	memory := make(map[int]int)

	mask := ""
	for _, command := range commands {
		m1 := maskRegex.FindStringSubmatch(command)
		if len(m1) > 1 {
			mask = m1[1]
			fmt.Printf("%v\n", mask)
			continue
		}

		m2 := memRegex.FindStringSubmatch(command)
		if len(m2) > 1 {
			dst, _ := strconv.Atoi(m2[1])
			val, _ := strconv.Atoi(m2[2])

			// Convert value.
			//fmt.Printf("from value %v\n", val)
			for i := len(mask) - 1; i >= 0; i-- {
				if mask[i] == '1' {
					val = setBit(val, len(mask)-i-1)
				} else if mask[i] == '0' {
					val = clearBit(val, len(mask)-i-1)
				}
			}
			//fmt.Printf("masked value %v\n", val)

			memory[dst] = val
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

func setBit(n int, pos int) int {
	n |= 1 << pos
	return n
}

func clearBit(n int, pos int) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}
