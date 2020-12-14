package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	maskRegex = regexp.MustCompile(`mask = (.*)`)
	memRegex  = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
)

func day14() {
	commands := readLines("input/14.txt")
	memory := make(map[int]int64)

	mask := ""
	for _, command := range commands {
		m1 := maskRegex.FindStringSubmatch(command)
		if len(m1) > 1 {
			mask = m1[1]
			fmt.Printf("%v %d\n", mask, len(mask))
			continue
		}

		m2 := memRegex.FindStringSubmatch(command)
		if len(m2) > 1 {
			dst, _ := strconv.Atoi(m2[1])
			v, _ := strconv.Atoi(m2[2])
			val := int64(v)

			// Convert value.
			fmt.Printf("\nfrom value %v\n", val)
			for i := len(mask) - 1; i >= 0; i-- {
				if mask[i] == '1' {
					val = setBit(val, int64(len(mask)-i-1))
				} else if mask[i] == '0' {
					val = clearBit(val, int64(len(mask)-i-1))
				}
			}
			fmt.Printf("masked value %v\n", val)

			memory[dst] = int64(val)
			fmt.Printf("%v <- %v\n", dst, val)
			continue
		}

		panic("Unknown command:" + command)
	}

	// Result
	sum := int64(0)
	for _, v := range memory {
		sum += v
		fmt.Printf("Summing up %d -> %d\n", v, sum)
	}
	fmt.Printf("Result: %d\n", sum)
}

func setBit(n int64, pos int64) int64 {
	n |= 1 << pos
	return n
}

func clearBit(n int64, pos int64) int64 {
	var mask int64
	mask = ^(1 << pos)
	n &= mask
	return n
}
