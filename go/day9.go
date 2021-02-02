package main

import "math"

func day9() {
	numbers := linesToNumbers(readLines("input/9.txt"))
	preamble := 25

	var targetNumber int
	// Find first number which is not the sum of the last <preamble> numbers.
	for i := preamble; i < len(numbers); i++ {
		part := numbers[i-preamble : i]
		if !checkXMAS(part, numbers[i]) {
			println(numbers[i])
			targetNumber = numbers[i]
			break
		}
	}

	// Find continuous set of numbers
	var min, max int
outerLoop:
	for i := 0; i < len(numbers); i++ {
		sum := 0
		for j := i; j < len(numbers); j++ {
			sum += numbers[j]
			if sum == targetNumber {
				min = i
				max = j
				break outerLoop
			}
			if sum > targetNumber {
				break
			}
		}
	}

	// Find minimum and maximum number in this range
	minTarget := math.MaxInt64
	maxTarget := 0
	for i := min; i < max; i++ {
		if numbers[i] < minTarget {
			minTarget = numbers[i]
		}
		if numbers[i] > maxTarget {
			maxTarget = numbers[i]
		}
	}

	println(minTarget + maxTarget)
}

func checkXMAS(part []int, num int) bool {
	for i := 0; i < len(part); i++ {
		for j := 0; j < len(part); j++ {
			if i == j {
				continue
			}

			if part[i]+part[j] == num {
				return true
			}
		}
	}

	return false
}
