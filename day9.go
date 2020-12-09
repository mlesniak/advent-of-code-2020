package main

func day9() {
	numbers := linesToNumbers(readLines("input/9.txt"))
	preamble := 25

	// Find first number which is not the sum of the last <preamble> numbers.
	for i := preamble; i < len(numbers); i++ {
		part := numbers[i-preamble : i]
		if !checkXMAS(part, numbers[i]) {
			println(numbers[i])
			break
		}
	}
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
